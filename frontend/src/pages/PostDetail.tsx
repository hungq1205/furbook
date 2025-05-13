import React from 'react';
import { useParams, Link } from 'react-router-dom';
import { ChevronLeft, Heart, MessageCircle, Share2, Users, AlertTriangle } from 'lucide-react';
import { motion } from 'framer-motion';
import Button from '../components/common/Button';
import Avatar from '../components/common/Avatar';
import Card from '../components/common/Card';
import MediaGallery from '../components/feed/MediaGallery';
import { posts } from '../data/mockData';
import { formatDistanceToNow } from '../utils/dateUtils';

const PostDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const post = posts.find(p => p.id === id);

  if (!post) {
    return (
      <div className="text-center py-10">
        <AlertTriangle size={48} className="text-warning-500 mx-auto mb-4" />
        <h2 className="text-xl font-medium text-gray-900 mb-2">Post not found</h2>
        <p className="text-gray-600 mb-4">The post you're looking for doesn't exist or has been removed.</p>
        <Link to="/">
          <Button variant="primary">Back to Feed</Button>
        </Link>
      </div>
    );
  }

  const likeCount = post.interactions.filter(i => i.type === 'like').length;
  const shareCount = post.interactions.filter(i => i.type === 'share').length;

  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      transition={{ duration: 0.5 }}
    >
      <Link to="/" className="inline-flex items-center text-gray-600 hover:text-primary-600 mb-4">
        <ChevronLeft size={20} className="mr-1" />
        <span>Back to Feed</span>
      </Link>
      
      <Card>
        <div className="p-6">
          <div className="flex items-center mb-6">
            <Link to={`/profile/${post.username}`}>
              <Avatar src={post.userAvatar} alt={post.displayName} size="lg" />
            </Link>
            <div className="ml-4">
              <Link to={`/profile/${post.username}`} className="font-medium text-lg text-gray-900 hover:underline">
                {post.displayName}
              </Link>
              <p className="text-sm text-gray-500">
                {formatDistanceToNow(new Date(post.createdAt))}
              </p>
            </div>
          </div>
          
          {(post.type === 'lost' || post.type === 'found') && (
            <>
              <div className="mb-4 inline-block px-3 py-1 rounded-full text-sm font-medium bg-error-100 text-error-700">
                {post.type === 'lost' ? 'Missing' : 'Found'}
              </div>
            </>
          )}
          
          <p className="text-gray-800 text-lg mb-6">{post.content}</p>
          
          {(post.type === 'lost' || post.type === 'found') && (
            <>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
                <div>
                  <h3 className="text-sm font-medium text-gray-500 uppercase mb-2">Last Seen Location</h3>
                  <p className="text-gray-800">{post.lastSeen?.address}</p>
                </div>
                <div>
                  <h3 className="text-sm font-medium text-gray-500 uppercase mb-2">Area</h3>
                  <p className="text-gray-800">{post.area?.address}</p>
                </div>
                <div>
                  <h3 className="text-sm font-medium text-gray-500 uppercase mb-2">Contact Information</h3>
                  <p className="text-gray-800">{post.contactInfo}</p>
                </div>
                <div>
                  <h3 className="text-sm font-medium text-gray-500 uppercase mb-2">Helpers</h3>
                  <div className="flex items-center">
                    <Users size={18} className="text-primary-500 mr-2" />
                    <p className="text-gray-800">{post.participants?.length || 0} people helping</p>
                  </div>
                </div>
              </div>
            </>
          )}
          
          {post.medias?.length > 0 && (
            <div className="mb-6">
              <MediaGallery media={post.medias} />
            </div>
          )}
          
          <div className="flex flex-wrap items-center justify-between border-t border-b border-gray-200 py-4 my-6">
            {post.type === 'blog' ? (
              <>
                <div className="flex space-x-6 mb-2 md:mb-0">
                  <button className="flex items-center text-gray-600 hover:text-primary-600 transition-colors">
                    <Heart size={20} className="mr-2" />
                    <span>{likeCount} Likes</span>
                  </button>
                  <button className="flex items-center text-gray-600 hover:text-primary-600 transition-colors">
                    <MessageCircle size={20} className="mr-2" />
                    <span>{post.commentNum} Comments</span>
                  </button>
                  <button className="flex items-center text-gray-600 hover:text-primary-600 transition-colors">
                    <Share2 size={20} className="mr-2" />
                    <span>{shareCount} Shares</span>
                  </button>
                </div>
              </>
            ) : (
              <>
                <Button
                  variant={post.isResolved ? 'ghost' : 'secondary'}
                  icon={<Users size={18} />}
                  disabled={post.isResolved}
                >
                  {post.isResolved ? 'Pet Found' : 'Help Find'}
                </Button>
                <Button variant="outline" icon={<Share2 size={18} />}>
                  Share
                </Button>
              </>
            )}
          </div>
          
          <div>
            <h3 className="font-medium text-gray-900 mb-4">Comments</h3>
            <div className="flex space-x-3 mb-6">
              <Avatar src={post.userAvatar} alt={post.displayName} size="md" />
              <div className="flex-1">
                <textarea
                  placeholder="Write a comment..."
                  className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-transparent resize-none"
                  rows={2}
                ></textarea>
                <div className="flex justify-end mt-2">
                  <Button variant="primary" size="sm">Post Comment</Button>
                </div>
              </div>
            </div>
            
            {post.commentNum > 0 ? (
              <div className="space-y-4">
                <div className="flex space-x-3">
                  <Link to={`/profile/${posts[1].username}`}>
                    <Avatar 
                      src={posts[1].userAvatar} 
                      alt={posts[1].displayName} 
                      size="md" 
                    />
                  </Link>
                  <div>
                    <div className="bg-gray-100 p-3 rounded-lg justify-items-start">
                      <Link to={`/profile/${posts[1].username}`} className="mb-1 hover:underline">
                        <div className="font-medium text-gray-900">{posts[1].displayName}</div>
                      </Link>
                      <p className="text-gray-800">
                        {post.type === 'lost' || post.type === 'found'
                          ? "I think I saw this pet around Maple Street yesterday! I'll keep an eye out and let you know if I see them again."
                          : "Your pet is adorable! What breed are they?"}
                      </p>
                    </div>
                    <div className="flex items-center mt-1 text-xs text-gray-500">
                      <button className="mr-3 hover:text-primary-600">Like</button>
                      <button className="mr-3 hover:text-primary-600">Reply</button>
                      <span>2 hours ago</span>
                    </div>
                  </div>
                </div>
              </div>
            ) : (
              <div className="text-center py-4 text-gray-500">
                No comments yet. Be the first to comment!
              </div>
            )}
          </div>
        </div>
      </Card>
    </motion.div>
  );
};

export default PostDetail;