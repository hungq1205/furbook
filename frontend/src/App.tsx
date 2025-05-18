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
import { authService } from './services/authService';

// ProtectedRoute component
const ProtectedRoute = ({ children }: { children: JSX.Element }) => {
  return authService.isAuthenticated() ? children : <Navigate to="/login" replace />;
};

function App() {
  return (
    <Router>
      <AnimatePresence mode="wait">
        <Routes>
          <Route path="/login" element={<Navigate to="/" replace />} />

          <Route
            path="/"
            element={
              <ProtectedRoute>
                <MainLayout />
              </ProtectedRoute>
            }
          >
            <Route index element={<Feed />} />
            <Route path="lost-pets" element={<LostPets />} />
            <Route path="messages" element={<Messages />} />
            <Route path="profile" element={<Profile />} />
            <Route path="profile/:id" element={<Profile />} />
            <Route path="post/:id" element={<PostDetail />} />
            <Route path="lost-pets/:id" element={<PostDetail />} />
          </Route>

          <Route
            path="/"
            element={
              <ProtectedRoute>
                <SimpleLayout />
              </ProtectedRoute>
            }
          >
            <Route path="create-lost-pet" element={<CreateLostPet />} />
          </Route>

          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </AnimatePresence>
    </Router>
  );
}

export default App;