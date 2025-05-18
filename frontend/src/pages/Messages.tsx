import React, { useEffect, useState } from 'react';
import { motion } from 'framer-motion';
import { Search, Send } from 'lucide-react';
import Avatar from '../components/common/Avatar';
import Card from '../components/common/Card';
// import { friends, messages, groupChats } from '../data/mockData';
import { formatDistanceToNow } from '../utils/dateUtils';
import { groupChatService } from '../services/groupChatService';
import { GroupChat, Message } from '../types/message';
import { handleError } from '../utils/errors';
import { messageService } from '../services/messageService';
import { authService } from '../services/authService';

const selectGroupChat = async (
  selectedGroupId: number, 
  setSelectedGroup: (group: GroupChat | null) => void,
  setMessages: (messages: Message[]) => void,
) => {
  try {
    const group = await groupChatService.getGroupDetails(selectedGroupId)
    const messages = await messageService.getGroupMessages(selectedGroupId);
    setSelectedGroup(group);
    setMessages(messages);
  } catch (error) {
    handleError(error, 'Failed to fetch group chat details');
  }
}

const Messages: React.FC = () => {
  const [groupChats, setGroupChats] = useState<GroupChat[]>([]);
  const [selectedGroup, setSelectedGroup] = useState<GroupChat | null>(null);
  const [messages, setMessages] = useState<Message[]>([]);
  const [message, setMessage] = useState('');

  const friends = authService.getCurrentUserFriends();

  useEffect(() => {
    groupChatService.getGroups()
      .then(setGroupChats)
      .catch(error => handleError(error, 'Failed to fetch group chats'));
  }, [])

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
      className="h-[calc(100vh-theme(spacing.16))] md:h-[calc(100vh-theme(spacing.12))] flex flex-col"
    >
      <Card className="flex-1 overflow-hidden">
        <div className="h-full flex">
          <div className={`${selectedGroup ? 'hidden md:block' : 'w-full'} md:w-80 border-r border-gray-200`}>
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
                {groupChats.map(group  => {
                  const chatFriend = group.isDirect 
                    ? friends.find(f => group.members.includes(f.username))
                    : null;

                  return (
                    <button
                      key={group.id}
                      onClick={() => selectGroupChat(group.id, setSelectedGroup, setMessages)}
                      className={`w-full flex items-center p-4 hover:bg-gray-50 transition-colors border-b border-gray-100 ${
                        selectedGroup?.id === group.id ? 'bg-primary-50' : ''
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
                            {group.isDirect ? chatFriend?.displayName : group.name}
                          </p>
                          {group.lastMessage && (
                            <span className="text-xs text-gray-500">
                              {formatDistanceToNow(new Date(group.lastMessage.createdAt))}
                            </span>
                          )}
                        </div>
                        {group.lastMessage && (
                          <p className="text-sm text-gray-500 truncate">
                            {group.lastMessage.content}
                          </p>
                        )}
                      </div>
                    </button>
                  );
                })}
              </div>
            </div>
          </div>

          {/* Chat area */}
          <div className={`${selectedGroup ? 'flex' : 'hidden md:flex'} flex-1 flex-col h-full bg-gray-50`}>
            {selectedGroup ? (
              <>
                <div className="p-4 bg-white border-b border-gray-200 flex items-center justify-between">
                  <div className="flex items-center">
                    <button 
                      className="md:hidden mr-2 text-gray-600"
                      onClick={() => {
                        setSelectedGroup(null);
                        setMessages([]);
                      }}
                    >
                      Back
                    </button>
                    <span className="font-medium">
                      {selectedGroup.isDirect 
                        ? friends.find(f => selectedGroup.members.includes(f.username))?.displayName
                        : selectedGroup.name}
                    </span>
                  </div>
                </div>

                <div className="flex-1 overflow-y-auto p-4">
                  {messages.length > 0 ? (
                    messages.map((msg) => {
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