import React, { useEffect, useState } from 'react';
import { motion } from 'framer-motion';
import { Plus } from 'lucide-react';
import { Link } from 'react-router-dom';
import LostPetCard from '../components/lost-pet/LostPetCard';
import Button from '../components/common/Button';
import { Post } from '../types/post';
import { postService } from '../services/postService';
import { handleError } from '../utils/errors';
import { useAuth } from '../services/authService';

const LostPets: React.FC = () => {
  const authService = useAuth();
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    postService.getByUsers(authService.currentUserFriends!.map(f => f.username))
      .then(posts => setPosts(posts.filter(post => post.type === 'lost' || post.type === 'found')))
      .catch(error => handleError(error, 'Failed to fetch posts', authService.logout));
  }, []);

  return (
    <div>
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
      >
        <div className="flex items-center justify-between mb-6">
          <h1 className="text-2xl font-bold text-gray-900">Lost & Found Pets</h1>
          <Link to="/create-lost-pet">
            <Button variant="primary" icon={<Plus size={16} />}>
              Report Pet
            </Button>
          </Link>
        </div>
        
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {posts.map((post, index) => (
            <motion.div
              key={post.id}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.3, delay: index * 0.1 }}
            >
              <LostPetCard post={post} />
            </motion.div>
          ))}
        </div>
      </motion.div>
    </div>
  );
};

export default LostPets;