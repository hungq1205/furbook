import React from 'react';
import { Outlet, Link } from 'react-router-dom';
import { PawPrint, Home } from 'lucide-react';

const SimpleLayout: React.FC = () => {
  return (
    <div className="min-h-screen bg-gray-50 flex flex-col">
      {/* Header */}
      <header className="bg-white border-b border-gray-200 py-4">
        <div className="max-w-7xl mx-auto px-4 flex justify-between items-center">
          <Link to="/" className="text-xl font-bold text-primary-600 flex items-center">
            <PawPrint className="mr-2" size={24} />
            <span>FurBook</span>
          </Link>
          
          <div 
            onClick={() => window.history.back()} 
            className="flex items-center text-gray-600 hover:text-primary-600 transition-colors cursor-pointer text-lg mr-3"
          >
            <span>Back</span>
          </div>
        </div>
      </header>
      
      {/* Main content area */}
      <main className="flex-1">
        <div className="max-w-4xl mx-auto px-4 py-6">
          <Outlet />
        </div>
      </main>
      
      {/* Footer */}
      <footer className="bg-white border-t border-gray-200 py-4">
        <div className="max-w-7xl mx-auto px-4 text-center text-sm text-gray-500">
          &copy; {new Date().getFullYear()} FurBook. All rights reserved.
        </div>
      </footer>
    </div>
  );
};

export default SimpleLayout;