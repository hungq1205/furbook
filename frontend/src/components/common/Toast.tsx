import React, { useEffect, useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { Link, useNavigate } from 'react-router-dom';
import { X } from 'lucide-react';
import { formatNotification, Notification } from '../../types/notification';
import wsService from '../../services/webSocketService';
import Avatar from './Avatar';

const Toast: React.FC = () => {
  const navigate = useNavigate();
  const [noti, setNoti] = useState<Notification | null>(null);

  useEffect(() => wsService.subscribe<Notification>('notification', noti => setNoti(formatNotification(noti))), []);

  useEffect(() => {
    if (noti == null) return;
    const timer = setTimeout(() => setNoti(null), 5000);
    return () => clearTimeout(timer);
  }, [noti]);

  const content = noti ? (
    <motion.div
      initial={{ opacity: 0, y: 50, x: '100%' }}
      animate={{ opacity: 1, y: 0, x: 0 }}
      exit={{ opacity: 0, y: 20, x: '100%' }}
      className="fixed bottom-4 right-4 flex items-center bg-white text-gray-800 rounded-lg shadow-md overflow-hidden max-w-sm z-50"
    >
      <div className="w-1 self-stretch bg-primary-500" />
      <div className="flex items-center p-3">
        <Avatar 
          src={noti.icon}
          alt='noti'
          size='md'
        />
        <p className="text-sm font-medium">{noti.desc}</p>
        <button
          onClick={(e) => {
            e.preventDefault();
            e.stopPropagation();
            setNoti(null);
            navigate('/' + noti.link)
          }}
          className="ml-4 p-1 hover:bg-gray-100 rounded-full transition-colors"
        >
          <X size={16} className="text-gray-500" />
        </button>
      </div>
    </motion.div>
  ) : null;

  return (
    <AnimatePresence>
      {
        noti && (
          noti.link ? 
            <Link to={noti.link}>{content}</Link> : 
            content
        )
      }
    </AnimatePresence>
  );
};

export default Toast;