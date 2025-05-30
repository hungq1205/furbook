import React, { useEffect, useRef, useState } from 'react';
import { motion } from 'framer-motion';
import { useParams } from 'react-router-dom';
import Avatar from '../components/common/Avatar';
import Button from '../components/common/Button';
import Card from '../components/common/Card';
// import { posts, currentUser, friends } from '../data/mockData';
import PostCard from '../components/feed/PostCard';
import LostPetCard from '../components/lost-pet/LostPetCard';
import { useAuth } from '../services/authService';
import { Post } from '../types/post';
import { Friendship, User } from '../types/user';
import { userService } from '../services/userService';
import { postService } from '../services/postService';
import { handleError } from '../utils/errors';
import { Edit } from 'lucide-react';
import { fileService } from '../services/fileService';

const Profile: React.FC = () => {
  const authService = useAuth();
  const { username } = useParams<{ username: string }>();
  const [profileUser, setProfileUser] = useState<User | null>();
  const [activeTab, setActiveTab] = useState<'posts' | 'lost-pets' | 'participated'>('posts');
  const [friendStatus, setFriendStatus] = useState<Friendship>('none');
  const [posts, setPosts] = useState<Post[]>([]); 
  const lostPosts = useRef<Post[]>([]);

  const currentUser = authService.currentUser!;

  const isOwnProfile = !username || username === currentUser.username;

  useEffect(() => {
    if (isOwnProfile) {
      setProfileUser(currentUser);
      return;
    }
    userService.getUser(username)
      .then(setProfileUser)
      .catch(error => {
        console.error('Failed to fetch user:', error);
        setProfileUser(null);
      });
    userService.checkFriendship(username)
      .then(res => setFriendStatus(res.friendship))
      .catch(error => console.error('Failed to check friendship:', error));
  }, [username]);

  useEffect(() => {
    if (!profileUser) return;
    postService.getByUsername(profileUser.username)
      .then(posts => {
        setPosts(posts);
        lostPosts.current = posts.filter(post => post.type !== 'blog');
      })
      .catch(error => console.error('Failed to fetch posts:', error));
  }, [profileUser]);
  
  const handleDelete = (id: string) => {
    if (!profileUser) return;
    lostPosts.current = lostPosts.current.filter(p => p.id !== id);
    setPosts(posts.filter(p => p.id !== id));
  }

  const handleFriendRequest = (resultFriendship: Friendship) => {
    if (!profileUser) return;
    userService.sendFriendRequest(profileUser.username)
      .then(() => {
        setFriendStatus(resultFriendship);
        authService.refresh();
      })
      .catch(error => handleError('Failed to sent friend request:', error, authService.logout));
  }

  const handleFriendRequestRevoke = (resultFriendship: Friendship) => {
    if (!profileUser) return;
    userService.revokeFriendRequest(profileUser.username)
      .then(() => setFriendStatus(resultFriendship))
      .catch(error => handleError('Failed to revoke friend request:', error, authService.logout));
  }

  const handleUnfriend = (resultFriendship: Friendship) => {
    if (!profileUser) return;
    userService.removeFriend(profileUser.username)
      .then(() => {
        setFriendStatus(resultFriendship);
        authService.refresh();
      })
      .catch(error => handleError('Failed to unfriend request:', error, authService.logout));
  }

  const handleFriendRequestDecline = () => {
    if (!profileUser) return;
    userService.declineFriendRequest(profileUser.username)
      .then(() => setFriendStatus('none'))
      .catch(error => handleError('Failed to decline friend request:', error, authService.logout));
  }

  const handleFriendAction = () => {
    switch (friendStatus) {
      case 'none':
        handleFriendRequest('sent');
        break;
      case 'sent':
        handleFriendRequestRevoke('none');
        break;
      case 'received':
        handleFriendRequest('friend');
        break;
      case 'friend':
        if (window.confirm(`Are you sure you want to unfriend with this ${profileUser?.displayName} #${profileUser?.username}?`))
          handleUnfriend('none');
        break;
    }
  };

  const handleUploadAvatar = async (e: React.ChangeEvent<HTMLInputElement>) => {
    if (!e.target.files || e.target.files.length < 0) 
      return;
    try {
      const url = await fileService.upload(e.target.files[0]);
      const updatedUser = await userService.updateUser({username: currentUser.username, avatar: url});
      await authService.refresh()
      if (profileUser?.username === currentUser.username)
        setProfileUser(updatedUser)
      e.target.value = '';
    } catch (err) {
      handleError(err, 'failed to upload avatar', authService.logout)
    }
  };

  const tabs = [
    { username: 'posts', label: 'Posts', count: posts.length },
    { username: 'lost-pets', label: 'Lost & Found', count: lostPosts.current.length },
    { username: 'participated', label: 'Participated', count: 1 },
  ];

  if (!profileUser)
    return;

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
    >
      <Card className="mb-6 overflow-hidden">
        <div className="h-48 bg-gradient-to-r from-primary-500 to-secondary-500"></div>
        <div className="p-6 relative">
            <div className="absolute top-0 transform -translate-y-1/2 left-6 cursor-pointer group">
              <div className="relative">
                <Avatar
                  src={profileUser.avatar}
                  alt={profileUser.displayName}
                  size="xl"
                />
                <div 
                  className="absolute text-white inset-0 border-2 rounded-full bg-black bg-opacity-30 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity"
                  onClick={() => document.getElementById('avatar-input')?.click()}
                >
                  <Edit size={30} />
                </div>
              </div>
              <input
                id="avatar-input"
                type="file"
                multiple
                accept="image/*"
                className="hidden"
                onChange={handleUploadAvatar}
              />
            </div>
          
          <div className="pt-12 flex items-start">
            <div>
              <h1 className="text-2xl font-bold text-gray-900">{profileUser.displayName}</h1>
              <p className="text-gray-600 mt-1">{profileUser.bio}</p>
              
              <div className="flex flex-wrap gap-6 mt-4">
                <div className="text-center">
                  <p className="text-xl font-bold text-gray-900">{profileUser.friendNum}</p>
                  <p className="text-sm text-gray-500">Friends</p>
                </div>
                <div className="text-center">
                  <p className="text-xl font-bold text-gray-900">{1}</p>
                  <p className="text-sm text-gray-500">Helped</p>
                </div>
              </div>
            </div>
            <div className="grow"/>
            {!isOwnProfile && (
              <>
              <Button
                variant={friendStatus === 'friend' ? 'outline' : (friendStatus === 'sent' ? 'warning' : 'primary')}
                ring={false}
                onClick={handleFriendAction}
              >
                {friendStatus === 'none' && 'Add Friend'}
                {friendStatus === 'sent' && 'Revoke Request'}
                {friendStatus === 'received' && 'Accept'}
                {friendStatus === 'friend' && 'Unfriend'}
              </Button>
              { friendStatus === 'received' && 
                <>
                <div className="w-2"/>
                <Button variant='warning' ring={false} onClick={handleFriendRequestDecline}> Decline </Button> 
                </>
              }
              </>
            )}
          </div>
        </div>
      </Card>
      
      <div className="mb-6">
        <div className="border-b border-gray-200">
          <nav className="-mb-px flex space-x-8">
            {tabs.map((tab) => (
              <button
                key={tab.username}
                className={`whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm ${
                  activeTab === tab.username
                    ? 'border-primary-500 text-primary-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
                onClick={() => setActiveTab(tab.username as typeof activeTab)}
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
            {posts.length > 0 ? (
              posts.map(post => (
                <motion.div
                  key={post.id}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.3 }}
                >
                  <PostCard post={post} onDelete={handleDelete}/>
                </motion.div>
              ))
            ) : (
              <p className="text-center py-8 text-gray-500">No posts yet.</p>
            )}
          </div>
        )}
        
        {activeTab === 'lost-pets' && (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {lostPosts.current.length > 0 ? (
              lostPosts.current.map(post => (
                <motion.div
                  key={post.id}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.3 }}
                >
                  <LostPetCard post={post} userLocation={undefined}/>
                </motion.div>
              ))
            ) : (
              <p className="text-center py-8 text-gray-500 col-span-2">No lost pet posts yet.</p>
            )}
          </div>
        )}
        
        {activeTab === 'participated' && (
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {posts.length > 0 ? (
              posts.map(post => (
                <motion.div
                  key={post.id}
                  initial={{ opacity: 0, y: 10 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.3 }}
                >
                  <LostPetCard post={post} userLocation={undefined}/>
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