import React, { useEffect, useState } from 'react';
import { motion } from 'framer-motion';
import { Link, useNavigate } from 'react-router-dom';
import Avatar from '../components/common/Avatar';
import Card from '../components/common/Card';
import { formatDistanceToNow } from '../utils/common';
import { formatNotification, Notification } from '../types/notification';
import { useAuth } from '../services/authService';
import { notiService } from '../services/notiService';
import { handleError } from '../utils/errors';

const Notifications: React.FC = () => {
  const authService = useAuth();
  const navigate = useNavigate();
  const [notifications, setNotifications] = useState<Notification[]>([]);
  const [page, setPage] = useState(1);

  const handleNotiClick = (id: string, url: string) => {
    notiService.markRead(id, true)
      .catch(err => handleError(err, 'failed to mark notification read', authService.logout))
      .finally(() => navigate('/' + url));
  };

  useEffect(() => {
    if (!authService.currentUser) return;
    notiService.getByUsername(page)
      .then(fn => {
        if (notifications.length > 0 && fn[0].id === notifications[0].id) return;
        setNotifications([...notifications, ...fn.map(formatNotification)].reverse());
      })
      .catch(err => handleError(err, 'failed to get notifications', authService.logout))
  }, [authService.currentUser, page])

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
    >
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold text-gray-900">Notifications</h1>
      </div>
      <Card>
        <div className="divide-y divide-gray-100">
          {notifications.map((notification) => (
            <div
              key={notification.id}
              onClick={() => handleNotiClick(notification.id, notification.link)}
              className={`block px-4 py-3 transition-colors hover:bg-gray-50 cursor-pointer ${
                notification.read ? 'bg-white' : 'bg-blue-50'
              }`}
            >
              <div className="flex items-start">
                <div className="flex-shrink-0 mr-3">
                  <Avatar
                    src={notification.icon}
                    alt="notification"
                    size="md"
                  />
                </div>
                
                <div className="flex-1 min-w-0">
                  <div className="flex items-center mb-1">
                    {notification.desc}
                  </div>
                  <div className="flex items-center text-xs text-gray-500">
                    {formatDistanceToNow(new Date(notification.created_at))}
                  </div>
                </div>
                
                {!notification.read && (
                  <div className="ml-3 flex-shrink-0">
                    <div className="w-2 h-2 rounded-full bg-primary-500"></div>
                  </div>
                )}
              </div>
            </div>
          ))}
        </div>
      </Card>
    </motion.div>
  );
};

export default Notifications;