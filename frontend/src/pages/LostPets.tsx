import React from 'react';
import { motion } from 'framer-motion';
import { Plus } from 'lucide-react';
import { Link } from 'react-router-dom';
import LostPetCard from '../components/lost-pet/LostPetCard';
import Button from '../components/common/Button';
import { posts } from '../data/mockData';

const LostPets: React.FC = () => {
  // Filter only lost and found pet posts
  const lostPetPosts = posts.filter(post => post.type === 'lost' || post.type === 'found');

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
          {lostPetPosts.map((post, index) => (
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