import React from 'react';
import { Outlet, useLocation } from 'react-router-dom';
import Sidebar from './Sidebar';
import FriendsSidebar from './FriendsSidebar';
import MobileNav from './MobileNav';

const MainLayout: React.FC = () => {
  const location = useLocation();
  const isMessagesRoute = location.pathname === '/messages';

  return (
    <div className="min-h-screen bg-gray-50 flex">
      <div className="w-64 hidden md:block bg-white border-r border-gray-200 min-h-screen">
        <Sidebar />
      </div>
      
      <div className="flex-1 min-w-0 pb-16 md:pb-0 relative">
        <main className={!isMessagesRoute ? 'max-w-5xl mx-auto px-4 py-6' : 'px-4 py-6'}>
          <Outlet />
        </main>
      </div>
      
      {!isMessagesRoute && (
        <div className="w-72 hidden lg:block bg-white border-l border-gray-200 min-h-screen">
          <FriendsSidebar />
        </div>
      )}

      <MobileNav />
    </div>
  );
};

export default MainLayout;