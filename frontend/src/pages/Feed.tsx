import React, { useEffect, useState } from 'react';
import { motion } from 'framer-motion';
import CreatePostInput from '../components/feed/CreatePostInput';
import PostCard from '../components/feed/PostCard';
import { Post } from '../types/post';
import { postService } from '../services/postService';
import { handleError } from '../utils/errors';
import { authService } from '../services/authService';
// import { posts } from '../data/mockData';

const Feed: React.FC = () => {
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    postService.getByUsers(authService.getCurrentUserFriends().map(f => f.username))
      .then(posts => setPosts(posts))
      .catch(error => handleError(error, 'Failed to fetch posts'));
  }, []);

  useEffect(() => {
    console.log('Posts:', posts);
  }, [posts]);
  
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
          {posts.map((post) => (
            <div
              key={post.id}
            >
              <PostCard post={post} />
            </div>
          ))}
        </div>
      </motion.div>
    </div>
  );
};

export default Feed;