import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { Search, Send } from 'lucide-react';
import Avatar from '../components/common/Avatar';
import Card from '../components/common/Card';
import { friends, messages, groupChats } from '../data/mockData';
import { formatDistanceToNow } from '../utils/dateUtils';

const Messages: React.FC = () => {
  const [selectedChat, setSelectedChat] = useState<number | null>(null);
  const [message, setMessage] = useState('');

  const selectedChatData = groupChats.find(chat => chat.id === selectedChat);
  const chatMessages = messages.filter(msg => msg.groupId === selectedChat);

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
      className="h-[calc(100vh-theme(spacing.16))] md:h-[calc(100vh-theme(spacing.12))] flex flex-col"
    >
      <Card className="flex-1 overflow-hidden">
        <div className="h-full flex">
          <div className={`${selectedChat ? 'hidden md:block' : 'w-full'} md:w-80 border-r border-gray-200`}>
            <div className="h-full flex flex-col">
              <div className="overflow-y-auto flex-1">
                <div className="relative m-4">
                  <input
                    type="text"
                    placeholder="Find a conversation..."
                    className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent"
                  />
                  <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400" size={20} />
                </div>
                {groupChats.map((chat) => {
                  const lastMessage = messages
                    .filter(m => m.groupId === chat.id)
                    .sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime())[0];
                  
                  const chatFriend = chat.isDirect 
                    ? friends.find(f => chat.members.includes(f.username))
                    : null;

                  return (
                    <button
                      key={chat.id}
                      onClick={() => setSelectedChat(chat.id)}
                      className={`w-full flex items-center p-4 hover:bg-gray-50 transition-colors border-b border-gray-100 ${
                        selectedChat === chat.id ? 'bg-primary-50' : ''
                      }`}
                    >
                      {chatFriend && (
                        <Avatar
                          src={chatFriend.avatar}
                          alt={chatFriend.displayName}
                          size="md"
                        />
                      )}
                      <div className="ml-3 flex-1 text-left">
                        <div className="flex items-center justify-between">
                          <p className="font-medium text-gray-900">
                            {chat.isDirect ? chatFriend?.displayName : chat.name}
                          </p>
                          {lastMessage && (
                            <span className="text-xs text-gray-500">
                              {formatDistanceToNow(new Date(lastMessage.createdAt))}
                            </span>
                          )}
                        </div>
                        <p className="text-sm text-gray-500 truncate">
                          {lastMessage?.content || 'No messages yet'}
                        </p>
                      </div>
                    </button>
                  );
                })}
              </div>
            </div>
          </div>

          {/* Chat area */}
          <div className={`${selectedChat ? 'flex' : 'hidden md:flex'} flex-1 flex-col h-full bg-gray-50`}>
            {selectedChatData ? (
              <>
                <div className="p-4 bg-white border-b border-gray-200 flex items-center justify-between">
                  <div className="flex items-center">
                    <button 
                      className="md:hidden mr-2 text-gray-600"
                      onClick={() => setSelectedChat(null)}
                    >
                      Back
                    </button>
                    <span className="font-medium">
                      {selectedChatData.isDirect 
                        ? friends.find(f => selectedChatData.members.includes(f.username))?.displayName
                        : selectedChatData.name}
                    </span>
                  </div>
                </div>

                <div className="flex-1 overflow-y-auto p-4">
                  {chatMessages.length > 0 ? (
                    chatMessages.map((msg) => {
                      const sender = friends.find(f => f.username === msg.username);
                      return (
                        <div key={msg.id} className="mb-4">
                          <div className="flex items-start">
                            <Avatar
                              src={sender?.avatar || ''}
                              alt={sender?.displayName || ''}
                              size="sm"
                            />
                            <div className="ml-2">
                              <div className="flex items-center">
                                <span className="font-medium text-sm">{sender?.displayName}</span>
                                <span className="text-xs text-gray-500 ml-2">
                                  {formatDistanceToNow(new Date(msg.createdAt))}
                                </span>
                              </div>
                              <p className="text-gray-800 mt-1">{msg.content}</p>
                            </div>
                          </div>
                        </div>
                      );
                    })
                  ) : (
                    <div className="text-center text-sm text-gray-500 py-2">
                      Start a conversation
                    </div>
                  )}
                </div>

                <div className="p-4 bg-white border-t border-gray-200">
                  <div className="flex space-x-2">
                    <input
                      type="text"
                      value={message}
                      onChange={(e) => setMessage(e.target.value)}
                      placeholder="Type a message..."
                      className="flex-1 px-4 py-2 border border-gray-300 rounded-full focus:outline-none focus:ring-2 focus:ring-primary-500"
                    />
                    <button
                      className="p-2 bg-primary-600 text-white rounded-full hover:bg-primary-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                      disabled={!message.trim()}
                    >
                      <Send size={20} />
                    </button>
                  </div>
                </div>
              </>
            ) : (
              <div className="flex-1 flex items-center justify-center">
                <div className="text-center">
                  <h3 className="text-lg font-medium text-gray-900 mb-2">
                    Select a Conversation
                  </h3>
                  <p className="text-gray-500">
                    Choose a chat from the list to start messaging
                  </p>
                </div>
              </div>
            )}
          </div>
        </div>
      </Card>
    </motion.div>
  );
};

export default Messages;