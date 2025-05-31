import React, { useEffect, useState } from 'react';
import { motion } from 'framer-motion';
import CreatePostInput from '../components/feed/CreatePostInput';
import PostCard from '../components/feed/PostCard';
import { Post } from '../types/post';
import { postService } from '../services/postService';
import { handleError } from '../utils/errors';
import { useAuth } from '../services/authService';
// import { posts } from '../data/mockData';

const Feed: React.FC = () => {
  const authService = useAuth();
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    if (!authService.currentUser) return;
    postService.getByUsers([...authService.currentUserFriends.map(f => f.username), authService.currentUser.username])
      .then(posts => setPosts(posts))
      .catch(error => handleError(error, 'Failed to fetch posts', authService.logout));
  }, [authService.currentUser]);
  
  const handleDelete = (id: string) => setPosts(posts.filter(p => p.id !== id))
  const handleUploaded = (post: Post) => setPosts([post, ...posts])

  return (
    <div>
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.5 }}
      >
        <CreatePostInput onUploaded={p => handleUploaded(p)}/>
        
        <div className="space-y-4">
          {posts.map((post) => (
            <div key={post.id}>
              <PostCard post={post} onDelete={handleDelete}/>
            </div>
          ))}
        </div>
      </motion.div>
    </div>
  );
};

export default Feed;