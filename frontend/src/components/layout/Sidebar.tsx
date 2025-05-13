import React from 'react';
import { NavLink } from 'react-router-dom';
import { Home, User, PawPrint, MessageCircle, Settings, LogOut } from 'lucide-react';
import { motion } from 'framer-motion';

interface SidebarProps {
  className?: string;
}

const Sidebar: React.FC<SidebarProps> = ({ className = '' }) => {
  const menuItems = [
    { icon: <Home size={24} />, label: 'Feed', path: '/' },
    { icon: <PawPrint size={24} />, label: 'Lost Pets', path: '/lost-pets' },
    { icon: <MessageCircle size={24} />, label: 'Messages', path: '/messages' },
    { icon: <User size={24} />, label: 'Profile', path: '/profile' },
  ];

  return (
    <aside className={`flex flex-col justify-between p-4 ${className}`}>
      <div>
        <div className="mb-8 pl-2">
          <h1 className="text-2xl font-bold text-primary-600 flex items-center">
            <PawPrint className="mr-2" size={28} />
            <span>FurBook</span>
          </h1>
        </div>
        
        <nav>
          <ul className="space-y-2">
            {menuItems.map((item) => (
              <li key={item.path}>
                <NavLink
                  to={item.path}
                  className={({ isActive }) =>
                    `flex items-center space-x-3 px-4 py-3 rounded-lg transition-colors ${
                      isActive
                        ? 'bg-primary-50 text-primary-600 font-medium'
                        : 'text-gray-700 hover:bg-gray-100'
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
                      <span>{item.label}</span>
                    </>
                  )}
                </NavLink>
              </li>
            ))}
          </ul>
        </nav>
      </div>
      
      <div className="mt-auto">
        <button className="w-full flex items-center space-x-3 px-4 py-3 rounded-lg text-gray-700 hover:bg-gray-100 transition-colors mt-2">
          <LogOut size={24} />
          <span>Logout</span>
        </button>
      </div>
    </aside>
  );
};

export default Sidebar;