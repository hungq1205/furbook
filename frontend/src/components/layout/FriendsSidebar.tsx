import React, { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { MessageCircle, X, Minus, Maximize2 } from 'lucide-react';
import Avatar from '../common/Avatar';
// import { friends } from '../../data/mockData';
import { authService } from '../../services/authService';

interface ChatTabProps {
  username: string;
  displayName: string;
  avatar: string;
  onClose: () => void;
  onMinimize: () => void;
  isMinimized: boolean;
}

const ChatTab: React.FC<ChatTabProps> = ({ displayName, avatar, onClose, onMinimize, isMinimized }) => {
  return (
    <motion.div
      initial={{ height: 0, opacity: 0 }}
      animate={{ height: 'auto', opacity: 1 }}
      exit={{ height: 0, opacity: 0 }}
      className="bg-white rounded-t-lg shadow-md overflow-hidden"
      style={{ width: '18rem' }}
    >
      <div className="p-3 bg-primary-600 flex items-center justify-between text-white">
        <div className="flex items-center space-x-2">
          <Avatar src={avatar} alt={displayName} size="sm" />
          <span className="font-medium">{displayName}</span>
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
          <div className="h-48 p-3 overflow-y-auto bg-gray-50">
            <div className="text-center text-sm text-gray-500 py-2">
              Start a conversation...
            </div>
          </div>
          
          <div className="p-2 border-t border-gray-200">
            <input
              type="text"
              placeholder="Type a message..."
              className="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-1 focus:ring-primary-500"
            />
          </div>
        </>
      )}
    </motion.div>
  );
};

const FriendsSidebar: React.FC = () => {
  const [openChats, setOpenChats] = useState<string[]>([]);
  const [minimizedChats, setMinimizedChats] = useState<string[]>([]);
  
  const toggleChat = (username: string) => {
    if (openChats.includes(username)) {
      if (minimizedChats.includes(username)) {
        setMinimizedChats(minimizedChats.filter(id => id !== username));
      } else {
        setMinimizedChats([...minimizedChats, username]);
      }
    } else {
      if (openChats.length >= 3) {
        setOpenChats([...openChats.slice(1), username]);
      } else {
        setOpenChats([...openChats, username]);
      }
    }
  };
  
  const closeChat = (username: string) => {
    setOpenChats(openChats.filter(id => id !== username));
    setMinimizedChats(minimizedChats.filter(id => id !== username));
  };

  const toggleMinimize = (username: string) => {
    if (minimizedChats.includes(username)) {
      setMinimizedChats(minimizedChats.filter(id => id !== username));
    } else {
      setMinimizedChats([...minimizedChats, username]);
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
        />
      </div>
      
      <div className="flex-1 overflow-y-auto p-2">
        <ul className="space-y-2">
          {authService.getCurrentUserFriends().map(friend => (
            <li key={friend.username}>
              <button
                onClick={() => toggleChat(friend.username)}
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
          {openChats.map(username => {
            const friend = authService.getCurrentUserFriends().find(f => f.username === username);
            if (!friend) return null;
            
            return (
              <motion.div
                className='self-end'
                key={username}
                initial={{ y: 20, opacity: 0 }}
                animate={{ y: 0, opacity: 1 }}
                exit={{ y: 20, opacity: 0 }}
              >
                <ChatTab
                  username={friend.username}
                  displayName={friend.displayName}
                  avatar={friend.avatar}
                  onClose={() => closeChat(friend.username)}
                  onMinimize={() => toggleMinimize(friend.username)}
                  isMinimized={minimizedChats.includes(friend.username)}
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