import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { Link, useParams } from 'react-router-dom';
import Avatar from '../components/common/Avatar';
import Button from '../components/common/Button';
import Card from '../components/common/Card';
import { posts, currentUser, friends } from '../data/mockData';
import PostCard from '../components/feed/PostCard';
import LostPetCard from '../components/lost-pet/LostPetCard';

const Profile: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [activeTab, setActiveTab] = useState<'posts' | 'lost-pets' | 'participated'>('posts');
  const [friendStatus, setFriendStatus] = useState<'none' | 'pending' | 'friends'>('none');
  
  const isOwnProfile = !id || id === currentUser.username;
  const profileUser = isOwnProfile ? currentUser : friends.find(f => f.username === id) || currentUser;
  
  const userPosts = posts.filter(post => post.username === profileUser.username && post.type === 'blog');
  const userLostPets = posts.filter(post => post.username === profileUser.username && (post.type === 'lost' || post.type === 'found'));
  const participatedPosts = posts.filter(post => 
    (post.type === 'lost' || post.type === 'found') && 
    post.participants?.includes(profileUser.username)
  );

  const handleFriendAction = () => {
    switch (friendStatus) {
      case 'none':
        setFriendStatus('pending');
        break;
      case 'pending':
        setFriendStatus('none');
        break;
      case 'friends':
        if (window.confirm('Are you sure you want to unfriend this user?')) {
          setFriendStatus('none');
        }
        break;
    }
  };
  
  const tabs = [
    { id: 'posts', label: 'Posts', count: userPosts.length },
    { id: 'lost-pets', label: 'Lost & Found', count: userLostPets.length },
    { id: 'participated', label: 'Participated', count: participatedPosts.length },
  ];

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
    >
      <Card className="mb-6 overflow-hidden">
        <div className="h-48 bg-gradient-to-r from-primary-500 to-secondary-500"></div>
        <div className="p-6 relative">
          <div className="absolute top-0 transform -translate-y-1/2 left-6">
            <Avatar src={profileUser.avatar} alt={profileUser.displayName} size="xl" />
          </div>
          
          <div className="pt-12 flex justify-between items-start">
            <div>
              <h1 className="text-2xl font-bold text-gray-900">{profileUser.displayName}</h1>
              <p className="text-gray-600 mt-1">{profileUser.bio}</p>
              
              <div className="flex flex-wrap gap-6 mt-4">
                <div className="text-center">
                  <p className="text-xl font-bold text-gray-900">{userPosts.length}</p>
                  <p className="text-sm text-gray-500">Posts</p>
                </div>
                <div className="text-center">
                  <p className="text-xl font-bold text-gray-900">{profileUser.friendNum}</p>
                  <p className="text-sm text-gray-500">Friends</p>
                </div>
                <div className="text-center">
                  <p className="text-xl font-bold text-gray-900">{participatedPosts.length}</p>
                  <p className="text-sm text-gray-500">Helped</p>
                </div>
              </div>
            </div>

            {!isOwnProfile && (
              <Button
                variant={friendStatus === 'friends' ? 'outline' : 'primary'}
                onClick={handleFriendAction}
              >
                {friendStatus === 'none' && 'Add Friend'}
                {friendStatus === 'pending' && 'Cancel Request'}
                {friendStatus === 'friends' && 'Unfriend'}
              </Button>
            )}
          </div>
        </div>
      </Card>
      
      <div className="mb-6">
        <div className="border-b border-gray-200">
          <nav className="-mb-px flex space-x-8">
            {tabs.map((tab) => (
              <button
                key={tab.id}
                className={`whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ${
                  activeTab === tab.id
                    ? 'border-primary-500 text-primary-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
                onClick={() => setActiveTab(tab.id as typeof activeTab)}
              >
                {tab.label}
                <span className="ml-2 rounded-full bg-gray-100 px-2.5 py-0.5 text-xs font-medium text-gray-800">
                  {tab.count}
                </span>
              </button>
            ))}
          </nav>
        </div>
      </div>
      
      <div>
        {activeTab === 'posts' && (
          <div className="space-y-4">
            {userPosts.length > 0 ? (
              userPosts.map(post => (
                <motion.div
                  key={post.id}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.3 }}
                >
                  <PostCard post={post} />
                </motion.div>
              ))
            ) : (
              <p className="text-center py-8 text-gray-500">No posts yet.</p>
            )}
          </div>
        )}
        
        {activeTab === 'lost-pets' && (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {userLostPets.length > 0 ? (
              userLostPets.map(post => (
                <motion.div
                  key={post.id}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.3 }}
                >
                  <LostPetCard post={post} />
                </motion.div>
              ))
            ) : (
              <p className="text-center py-8 text-gray-500 col-span-2">No lost pet posts yet.</p>
            )}
          </div>
        )}
        
        {activeTab === 'participated' && (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {participatedPosts.length > 0 ? (
              participatedPosts.map(post => (
                <motion.div
                  key={post.id}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.3 }}
                >
                  <LostPetCard post={post} />
                </motion.div>
              ))
            ) : (
              <p className="text-center py-8 text-gray-500 col-span-2">No participated searches yet.</p>
            )}
          </div>
        )}
      </div>
    </motion.div>
  );
};

export default Profile;