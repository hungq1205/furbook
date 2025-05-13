import React from 'react';
import { NavLink } from 'react-router-dom';
import { Home, User, PawPrint, MessageCircle } from 'lucide-react';
import { motion } from 'framer-motion';

const MobileNav: React.FC = () => {
  const menuItems = [
    { icon: <Home size={24} />, label: 'Feed', path: '/' },
    { icon: <PawPrint size={24} />, label: 'Lost Pets', path: '/lost-pets' },
    { icon: <MessageCircle size={24} />, label: 'Messages', path: '/messages' },
    { icon: <User size={24} />, label: 'Profile', path: '/profile' },
  ];

  return (
    <nav className="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 md:hidden">
      <div className="flex justify-around items-center">
        {menuItems.map((item) => (
          <NavLink
            key={item.path}
            to={item.path}
            className={({ isActive }) =>
              `flex flex-col items-center py-2 px-3 ${
                isActive ? 'text-primary-600' : 'text-gray-600'
              }`
            }
          >
            {({ isActive }) => (
              <>
                <motion.div
                  animate={{
                    scale: isActive ? 1.1 : 1,
                  }}
                  transition={{ duration: 0.2 }}
                >
                  {item.icon}
                </motion.div>
                <span className="text-xs mt-1">{item.label}</span>
              </>
            )}
          </NavLink>
        ))}
      </div>
    </nav>
  );
};

export default MobileNav;