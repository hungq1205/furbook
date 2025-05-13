import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { AnimatePresence } from 'framer-motion';

// Layouts
import MainLayout from './components/layout/MainLayout';
import SimpleLayout from './components/layout/SimpleLayout';

// Pages
import Feed from './pages/Feed';
import LostPets from './pages/LostPets';
import PostDetail from './pages/PostDetail';
import Profile from './pages/Profile';
import CreateLostPet from './pages/CreateLostPet';
import Messages from './pages/Messages';

function App() {
  return (
    <Router>
      <AnimatePresence mode="wait">
        <Routes>
          {/* Main layout routes */}
          <Route path="/" element={<MainLayout />}>
            <Route index element={<Feed />} />
            <Route path="lost-pets" element={<LostPets />} />
            <Route path="messages" element={<Messages />} />
            <Route path="profile" element={<Profile />} />
            <Route path="profile/:id" element={<Profile />} />
            <Route path="post/:id" element={<PostDetail />} />
            <Route path="lost-pets/:id" element={<PostDetail />} />
          </Route>
          
          {/* Simple layout routes */}
          <Route path="/" element={<SimpleLayout />}>
            <Route path="create-lost-pet" element={<CreateLostPet />} />
          </Route>
          
          {/* Redirect any other route to home */}
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </AnimatePresence>
    </Router>
  );
}

export default App;