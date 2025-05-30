import React, { useEffect, useRef, useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { MessageCircle, X, Minus, Maximize2 } from 'lucide-react';
import Avatar from '../common/Avatar';
// import { friends } from '../../data/mockData';
import { useAuth } from '../../services/authService';
import { messageService } from '../../services/messageService';
import { GroupChat, Message } from '../../types/message';
import { groupChatService } from '../../services/groupChatService';
import { Link, useNavigate } from 'react-router-dom';

interface ChatTabProps {
  groupId: number;
  username?: string;
  displayName: string;
  avatar?: string;
  onClose: () => void;
  onMinimize: () => void;
  isMinimized: boolean;
}

const ChatTab: React.FC<ChatTabProps> = ({ groupId, username, displayName, avatar, onClose, onMinimize, isMinimized }) => {
  const authService = useAuth();

  const [messages, setMessages] = useState<Message[]>([]);
  const [page, setPage] = useState(1);
  const isOutOfMessages = useRef(false);
  const chatContainerRef = useRef<HTMLDivElement>(null);
  const isFetchingMessages = useRef(false);

  const sendMessage = (groupId: number, content: string) => {
    if (!content.trim()) return;
    messageService.sendGroupMessage(groupId, content).then(newMessage => {
      setMessages(prevMessages => [...prevMessages, newMessage]);
      if (chatContainerRef.current)
        chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
    }).catch(error => {
      console.error('Error sending message:', error);
    });
  }

  const handleChatScroll = (e: React.UIEvent<HTMLDivElement>) => {
    const target = e.currentTarget as HTMLDivElement;
    if (target.scrollTop === 0 && !isFetchingMessages.current && !isOutOfMessages.current) {
      setPage((prevPage) => prevPage + 1);
    }
  }

  useEffect(() => {
    if (isOutOfMessages.current) return;
    messageService.getGroupMessages(groupId, page, 4).then(fetchedMessages => {
      isOutOfMessages.current = fetchedMessages.length == 0;
      isFetchingMessages.current = true;
      setMessages(prevMessages => {
        if (fetchedMessages[0]?.id !== prevMessages[0]?.id)
          return [...fetchedMessages, ...prevMessages]
        else
          return fetchedMessages;
      });
    }).catch(error => {
      console.error('Error fetching messages:', error);
    }).finally(() => isFetchingMessages.current = false);
  }, [page]);

  useEffect(() => {
    const chatContainer = chatContainerRef.current;
    if (chatContainer && page === 1)
      chatContainer.scrollTop = chatContainer.scrollHeight;
  }, [messages]);

  return (
    <motion.div
      initial={{ height: 0, opacity: 0 }}
      animate={{ height: 'auto', opacity: 1 }}
      exit={{ height: 0, opacity: 0 }}
      className="bg-white rounded-t-lg shadow-md overflow-hidden"
      style={{ width: '21rem' }}
    >
      <div className="p-3 bg-primary-600 flex items-center justify-between text-white">
        <div className="flex items-center space-x-2">
          { username ? 
            <>
              <Link to={`/profile/${username}`}>
                <Avatar src={avatar!} alt={displayName} size="sm" />
              </Link> 
              <Link to={`/profile/${username}`} className="font-medium hover:underline">
                {displayName}
              </Link>
            </> :
            <>
              <Avatar src={avatar!} alt={displayName} size="sm" />
              <span className="font-medium">{displayName}</span>
            </>
          }
        </div>
        <div className="flex items-center space-x-1">
          <button 
            onClick={onMinimize} 
            className="text-white hover:bg-primary-700 rounded-full p-1"
          >
            {isMinimized ? <Maximize2 size={16} /> : <Minus size={16} />}
          </button>
          <button 
            onClick={onClose} 
            className="text-white hover:bg-primary-700 rounded-full p-1"
          >
            <X size={16} />
          </button>
        </div>
      </div>
      
      {!isMinimized && (
        <>
            <div
              ref={chatContainerRef}
              className="h-64 p-3 overflow-y-auto bg-gray-50"
              onScroll={handleChatScroll}
            >
            {messages.length > 0 ? (
              messages.map((message, index) => (
                authService.currentUser!.username == message.username ? 
                <div key={index} className='mb-2 flex justify-end'>
                  <div className='flex items-end max-w-[85%] flex-row-reverse'>
                    <div
                      className='p-2 rounded-lg bg-primary-500 text-white'>
                      <div className="text-sm">{message.content}</div>
                    </div>
                    <div className="text-xs text-gray-500 mr-2 mb-1 w-8">
                        {new Date(message.created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', hour12: false })}
                    </div>
                  </div>
                </div> 
                :
                <div key={index} className='mb-2 flex justify-start'>
                  <div className="flex items-end max-w-[85%]">
                    <div
                      className='p-2 rounded-lg bg-gray-200 text-gray-900'
                    >
                      <div className="text-sm">{message.content}</div>
                    </div>
                    <div className="text-xs text-gray-500 ml-2 mb-1 w-6">
                    {new Date(message.created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', hour12: false })}
                    </div>
                  </div>
                </div>
              ))
            ) : (
              <div className="text-center text-sm text-gray-500 py-2">
                No messages yet. Start a conversation...
              </div>
            )}
            </div>
          
          <div className="p-2 border-t border-gray-200">
            <input
              type="text"
              placeholder="Type a message..."
              className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-1 focus:ring-primary-500"
              onKeyDown={(e) => {
                if (e.key === 'Enter' && !e.shiftKey) {
                  e.preventDefault();
                  const input = e.target as HTMLInputElement;
                  const content = input.value.trim();
                  if (content) {
                    sendMessage(groupId, content);
                    input.value = '';
                  }
                }
              }}
            />
          </div>
        </>
      )}
    </motion.div>
  );
};

const FriendsSidebar: React.FC = () => {
  const authService = useAuth();
  const navigate = useNavigate();

  const [query, setQuery] = useState("");
  const [chats, setChats] = useState<Map<number, GroupChat>>(new Map());
  const [openChats, setOpenChats] = useState<number[]>([]);
  const [minimizedChats, setMinimizedChats] = useState<number[]>([]);
  
  const fetchChat = async (id: number) => {
    try {
      const chat = await groupChatService.getGroupDetails(id);
      setChats(prevChats => new Map(prevChats).set(id, chat));
    } catch (error) {
      console.error('Error fetching chat:', error);
    }
  };

  const openChat = async (id: number) => {
    if (openChats.includes(id)) {
      setMinimizedChats(minimizedChats.filter(gid => gid !== id));
      return;
    }
    if (chats.has(id)) {
      setOpenChats([...openChats, id]);
      return;
    }
    await fetchChat(id);
    if (openChats.length >= 3) {
      setOpenChats([...openChats.slice(1), id]);
    } else {
      setOpenChats([...openChats, id]);
    }
  };
  
  const closeChat = (id: number) => {
    setOpenChats(openChats.filter(gid => gid !== id));
    setMinimizedChats(minimizedChats.filter(gid => gid !== id));
  };

  const toggleMinimize = (id: number) => {
    if (minimizedChats.includes(id)) {
      setMinimizedChats(minimizedChats.filter(gid => gid !== id));
    } else {
      setMinimizedChats([...minimizedChats, id]);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
      if (e.key === 'Enter') {
        e.preventDefault();
        navigate(`/profile/${e.currentTarget.value}`);
        setQuery('')
      }
    };

  return (
    <div className="flex flex-col h-full">
      <div className="p-4 border-b border-gray-200">
        <h3 className="font-medium text-gray-900 mb-2">Friends</h3>
        <input
          type="text"
          placeholder="Search friends..."
          className="w-full px-3 py-2 bg-gray-100 rounded-md text-sm focus:outline-none focus:bg-white focus:ring-1 focus:ring-primary-500"
          value={query}
          onKeyDown={handleKeyDown}
          onChange={e => setQuery(e.target.value)}
        />
        {query && (
          <Link to={`/profile/${query}`} className="text-sm text-gray-500 mt-2 hover:underline">
            Go to user @{query}
          </Link>
        )}
      </div>
      
      <div className="flex-1 overflow-y-auto p-2">
        <ul className="space-y-2">
          {authService.currentUserFriends
          .filter(friend => friend.displayName.toLowerCase().includes(query.toLowerCase()))
          .map(friend => (
            <li key={friend.username}>
              <button
                onClick={() => openChat(friend.groupid!)}
                className="w-full flex items-center p-2 rounded-md hover:bg-gray-100 transition-colors"
              >
                <Avatar src={friend.avatar} alt={friend.displayName} size="sm" />
                <div className="ml-3 text-left">
                  <p className="text-sm font-medium text-gray-900">{friend.displayName}</p>
                  <p className="text-xs text-gray-500">
                    {friend.bio}
                  </p>
                </div>
                <MessageCircle size={16} className="ml-auto text-gray-400" />
              </button>
            </li>
          ))}
        </ul>
      </div>
      
      <div className="fixed bottom-0 right-4 flex justify-end space-x-2 z-50">
        <AnimatePresence>
          {openChats.map(chatId => {
            const chat = chats.get(chatId);
            if (!chat) return null;
            return (
              <motion.div
                className='self-end'
                key={chat.id}
                initial={{ y: 20, opacity: 0 }}
                animate={{ y: 0, opacity: 1 }}
                exit={{ y: 20, opacity: 0 }}
              >
                <ChatTab
                  groupId={chat.id}
                  username={authService.currentUserFriends.find(f => f.groupid === chat.id)?.username}
                  displayName={chat.name}
                  avatar={chat.avatar}
                  onClose={() => closeChat(chat.id)}
                  onMinimize={() => toggleMinimize(chat.id)}
                  isMinimized={minimizedChats.includes(chat.id)}
                />
              </motion.div>
            );
          })}
        </AnimatePresence>
      </div>
    </div>
  );
};

export default FriendsSidebar;