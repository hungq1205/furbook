import React, { useEffect, useRef, useState } from 'react';
import { motion } from 'framer-motion';
import { Search, Send } from 'lucide-react';
import Avatar from '../components/common/Avatar';
import Card from '../components/common/Card';
// import { friends, messages, groupChats, currentUser } from '../data/mockData';
import { formatDistanceToNow } from '../utils/dateUtils';
import { groupChatService } from '../services/groupChatService';
import { GroupChat, Message } from '../types/message';
import { handleError } from '../utils/errors';
import { messageService } from '../services/messageService';
import { useAuth } from '../services/authService';
import { User } from '../types/user';
import { userService } from '../services/userService';

const Messages: React.FC = () => {
  const authService = useAuth();

  const [groupChats, setGroupChats] = useState<GroupChat[]>([]);
  const [selectedGroup, setSelectedGroup] = useState<GroupChat | null>(null);
  const [messages, setMessages] = useState<Message[]>([]);
  const [message, setMessage] = useState('');
  const [page, setPage] = useState(0);
  const [userDict, setUserDict] = useState<Map<string, User>>(new Map());
  const chatContainerRef = useRef<HTMLDivElement>(null);
  const isOutOfMessages = useRef(false);
  const isFetchingMessages = useRef(false);

  const currentUser = authService.currentUser!;

const updateUserDict = (...usernames: string[]) => {
  const missed = usernames.filter(username => !userDict.has(username));
  missed.length !== 0 && userService.getUsers(missed)
    .then(users => {
      const newDict = new Map(userDict);
      users.forEach(user => newDict.set(user.username, user));
      setUserDict(newDict);
    })
    .catch(err => handleError(err, 'Failed to fetch member details', authService.logout));
};

  const handleSelectChat = async (groupId: number) => {
    try {
      if (selectedGroup?.id !== groupId) {
        setMessages([]);
        const group = await groupChatService.getGroupDetails(groupId)
        isFetchingMessages.current = false;
        isOutOfMessages.current = false;

        if (!group.is_direct) updateUserDict(...group.members);
        setSelectedGroup(group);
        setPage(1);
      }
    } catch (error) {
      handleError(error, 'Failed to fetch group chat details', authService.logout);
    }
  };

  useEffect(() => {
    if (isOutOfMessages.current || !selectedGroup || page === 0) return;
    const fetchingChat = selectedGroup.id;
    messageService.getGroupMessages(selectedGroup.id, page, 4)
      .then(fetchedMessages => {
        if (selectedGroup?.id !== fetchingChat) return;

        isOutOfMessages.current = fetchedMessages.length === 0;
        isFetchingMessages.current = true;

        setMessages(prevMessages => {
          if (fetchedMessages[0]?.id !== prevMessages[0]?.id)
            return [...fetchedMessages, ...prevMessages]
          else
            return fetchedMessages;
        });
      })
      .catch(error => console.error('Error fetching messages:', error))
      .finally(() => isFetchingMessages.current = false);
  }, [selectedGroup?.id, page]);

  const handleSendMessage = async () => {
    if (!selectedGroup || !message.trim()) return;
    try {
      const newMessage = await messageService.sendGroupMessage(selectedGroup.id, message);
      setMessages(prevMessages => [...prevMessages, newMessage]);
      setMessage('');
    } catch (error) {
      handleError(error, 'Failed to send message', authService.logout);
    }
  };

  const handleChatScroll = (e: React.UIEvent<HTMLDivElement>) => {
    const target = e.currentTarget as HTMLDivElement;
    if (target.scrollTop === 0 && !isFetchingMessages.current && !isOutOfMessages.current)
      setPage((prevPage) => prevPage + 1);
  }

  useEffect(() => {
    const chatContainer = chatContainerRef.current;
    if (page === 1 && chatContainer)
      chatContainer.scrollTop = chatContainer.scrollHeight;
  }, [messages]);

  useEffect(() => {
    groupChatService.getGroups()
      .then(setGroupChats)
      .catch(error => handleError(error, 'Failed to fetch group chats', authService.logout));
  }, [])

  const MessageItem = ({ msg }: { msg: Message }) => (
    <div key={msg.id} className="mb-4">
      <div className="flex items-start">
        <Avatar
          src={userDict.get(msg.username)?.avatar ?? ''}
          alt={userDict.get(msg.username)?.displayName ?? ''}
          size="sm"
        />
        <div className="ml-2 max-w-[80%]">
          <div className="flex items-center">
            <span className="font-medium text-sm">{userDict.get(msg.username)?.displayName ?? ''}</span>
            <span className="text-xs text-gray-500 ml-2">
              {formatDistanceToNow(new Date(msg.created_at))}
            </span>
          </div>
          <p className="text-gray-800 mt-1">{msg.content}</p>
        </div>
      </div>
    </div>
  )

  const UserMessageItem = ({ msg }: { msg: Message }) => (
    <div key={msg.id} className="mb-4">
      <div className="flex items-start flex-row-reverse">
        <div className="mr-2 max-w-[80%] flex flex-col">
          <div className="flex items-center flex-row-reverse text-xs text-gray-500 mr-4">
              {formatDistanceToNow(new Date(msg.created_at))}
          </div>
          <p className="text-gray-800 mt-1 self-end">{msg.content}</p>
        </div>
      </div>
    </div>
  )

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
      className="h-[calc(100vh-theme(spacing.16))] md:h-[calc(100vh-theme(spacing.12))] flex flex-col"
    >
      <Card className="flex-1 overflow-hidden">
        <div className="h-full flex" onScroll={handleChatScroll}>
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
                {groupChats.map(group => {
                  return (
                    <button
                      key={group.id}
                      onClick={() => handleSelectChat(group.id)}
                      className={`w-full flex items-center p-4 hover:bg-gray-50 transition-colors border-b border-gray-100 ${
                        selectedGroup?.id === group.id ? 'bg-primary-50' : ''
                      }`}
                    >
                      {group.is_direct && (
                        <Avatar
                          src={group.avatar ?? ''}
                          alt={group.name}
                          size="md"
                        />
                      )}
                      <div className="ml-3 flex-1 text-left overflow-hidden">
                        <div className="flex items-center justify-between">
                          <p className="font-medium text-gray-900">
                            {group.name}
                          </p>
                          {group.last_message && (
                            <span className="text-xs text-gray-500">
                              {formatDistanceToNow(new Date(group.last_message.created_at))}
                            </span>
                          )}
                        </div>
                        {group.last_message && (
                          <p className="text-sm text-gray-500 truncate">
                            <b>{group.last_message.username + ': '}</b>{group.last_message.content}
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
                      {selectedGroup.name}
                    </span>
                  </div>
                </div>

                <div className="flex-1 overflow-y-auto p-4" ref={chatContainerRef}>
                  {messages.length > 0 ? (
                    messages.map((msg) => {
                      return ( msg.username === currentUser.username ?
                        <UserMessageItem key={msg.id} msg={msg} /> :
                        <MessageItem key={msg.id} msg={msg} />
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
                      onKeyDown={(e) => {
                        if (e.key === 'Enter') handleSendMessage();
                      }}
                      placeholder="Type a message..."
                      className="flex-1 px-4 py-2 border border-gray-300 rounded-full focus:outline-none focus:ring-2 focus:ring-primary-500"
                    />
                    <button
                      className="p-2 bg-primary-600 text-white rounded-full hover:bg-primary-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed w-10 h-10 flex items-center justify-center"
                      disabled={!message.trim()}
                    >
                      <Send 
                        className='pt-0.5 pr-0.5'
                        size={20} 
                        onClick={handleSendMessage}
                      />
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