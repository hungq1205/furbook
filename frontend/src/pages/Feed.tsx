import React from 'react';
import { motion } from 'framer-motion';
import CreatePostInput from '../components/feed/CreatePostInput';
import PostCard from '../components/feed/PostCard';
import { posts } from '../data/mockData';

const Feed: React.FC = () => {
  // Filter only blog posts
  const blogPosts = posts.filter(post => post.type === 'blog');

  return (
    <div>
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
      >
        <h1 className="text-2xl font-bold text-gray-900 mb-6">Your Feed</h1>
        <CreatePostInput />
        
        <div className="space-y-4">
          {blogPosts.map((post) => (
            <motion.div
              key={post.id}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.3, delay: 0.1 }}
            >
              <PostCard post={post} />
            </motion.div>
          ))}
        </div>
      </motion.div>
    </div>
  );
};

export default Feed;