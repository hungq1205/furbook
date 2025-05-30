import React, { useEffect, useMemo, useRef, useState } from 'react';
import { Heart, MessageCircle, Share2, MoreHorizontal, Edit, Trash2, Users } from 'lucide-react';
import { Link, useNavigate } from 'react-router-dom';
import { Post } from '../../types/post';
import Avatar from '../common/Avatar';
import Card from '../common/Card';
import IconButton from '../common/IconButton';
import MediaGallery from './MediaGallery';
import EditPostModal from '../post/EditPostModal';
import { BlogPostPayload } from '../../services/postService';
import { calcDistance, formatDistance, formatDistanceToNow } from '../../utils/common';
// import { currentUser } from '../../data/mockData';
import { postService } from '../../services/postService';
import { useAuth } from '../../services/authService';
import { handleError } from '../../utils/errors';
import { getTagColor } from '../lost-pet/LostPetCard';

interface PostCardProps {
  post: Post;
  onDelete: (id: string) => void;
}

const PostCard: React.FC<PostCardProps> = ({ post, onDelete }) => {
  const authService = useAuth();
  const navigate = useNavigate();

  const [userLocation, setUserLocation] = useState<{ lat: number; lng: number } | undefined>();
  const [showOptions, setShowOptions] = useState(false);
  const [showEditModal, setShowEditModal] = useState(false);
  const [userInteracted, setUserInteracted] = React.useState<Boolean>(false);
  const [interactionCount, setInteractionCount] = useState<number>(0);
  const isOwnPost = post.username === authService.currentUser!.username;

  const distance = useMemo(() => {
    if (!userLocation || !post.lastSeen) return "0";
    return formatDistance(calcDistance(
      userLocation.lat, userLocation.lng,
      post.lastSeen.lat, post.lastSeen.lng
    ));
  }, [userLocation, post.lastSeen]);

  useEffect(() => {
    if (navigator.geolocation)
      navigator.geolocation.getCurrentPosition(
        pos => setUserLocation({
          lat: pos.coords.latitude,
          lng: pos.coords.longitude
        }),
        err => console.error('Error getting location:', err)
      );
  }, []);

  useEffect(() => {
    const interacted = post.interactions.some(i => i.username === authService.currentUser!.username)
    const ic = interacted ? 
      post.interactions.length - 1 : 
      post.interactions.length;
    setUserInteracted(interacted);
    setInteractionCount(ic);
  }, [post]);

  const handleContentClick = () => {
    navigate(`/post/${post.id}`);
  };

  const handleOptionsClick = () => {
    setShowOptions(!showOptions);
  };

  const handleEdit = (e: React.MouseEvent) => {
    e.stopPropagation();
    setShowEditModal(true);
    setShowOptions(false);
  };

  const handleDelete = async (e: React.MouseEvent) => {
    e.stopPropagation();
    if (window.confirm('Are you sure you want to delete this post?')) {
      try {
        await postService.delete(post.id);
        onDelete(post.id)
      } catch (error) {
        console.error('Failed to delete post:', error);
      }
    }
    setShowOptions(false);
  };

  const handleSaveEdit = async () => {
    postService.updateContent(post.id, { content: post.content, medias: post.medias } as BlogPostPayload);
    setShowEditModal(false);
  };

  const handleInteract = () => {
    if (!post) return;
    if (userInteracted)
      postService.deleteInteraction(post.id)
        .then(() => userInteracted && setUserInteracted(false))
        .catch(error => handleError(error, 'Failed to unlike post', authService.logout));
    else
      postService.upsertInteraction(post.id)
        .then(() => !userInteracted && setUserInteracted(true))
        .catch(error => handleError(error, 'Failed to like post', authService.logout));
    setUserInteracted(!userInteracted);
  };

  const LostPostTags: React.FC = () => (
    post.type !== 'blog' && (
      <div className='mr-3'>
        { !post.isResolved && userLocation &&
          <div className="inline-block px-3 py-1 rounded-full text-sm font-medium bg-sky-100 text-sky-700">
          {distance} away
          </div>
        }
        <div className={`inline-block px-3 py-1 ml-2 rounded-full text-sm font-medium ${getTagColor(post.type, post.isResolved)}`}>
          {post.type === 'lost' ? 'Missing' : 'Found'}
        </div>
        { post.isResolved &&
          <div className="inline-block px-3 py-1 ml-2 rounded-full text-sm font-medium bg-success-100 text-success-700">
            Resolved
          </div>
        }
      </div>
    )
  );

  return (
    <>
      <Card className="mb-4">
        <div className="p-4">
          <div className="flex items-center justify-between mb-3">
            <div className="flex items-center space-x-3 w-full">
              <Link to={`/profile/${post.username}`}>
                <Avatar src={post.userAvatar} alt={post.displayName} size="md" />
              </Link>
              <div>
                <Link to={`/profile/${post.username}`} className="font-medium text-gray-900 hover:underline">
                  {post.displayName}
                </Link>
                <p className="text-xs text-gray-500">
                  {formatDistanceToNow(new Date(post.createdAt))}
                </p>
              </div>
              <div className='grow'/>
              { post.type !== 'blog' && (<div className='flex'><LostPostTags /></div>) }
            </div>
            {isOwnPost && (
              <div className="relative">
                <IconButton 
                  icon={<MoreHorizontal size={18} />} 
                  label="More options"
                  onClick={handleOptionsClick}
                />
                {showOptions && (
                  <div className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-10">
                    <button
                      className="w-full px-4 py-2 text-left text-sm text-gray-700 hover:bg-gray-100 flex items-center"
                      onClick={handleEdit}
                    >
                      <Edit size={16} className="mr-2" />
                      Edit Post
                    </button>
                    <button
                      className="w-full px-4 py-2 text-left text-sm text-red-600 hover:bg-gray-100 flex items-center"
                      onClick={handleDelete}
                    >
                      <Trash2 size={16} className="mr-2" />
                      Delete Post
                    </button>
                  </div>
                )}
              </div>
            )}
          </div>

          <p onClick={handleContentClick} className="text-gray-800 pb-3 cursor-pointer">{post.content}</p>
            
          {post.type === 'lost' && (
            <div className="cursor-pointer pb-6 pt-3" onClick={handleContentClick}>
              <div className="grid grid-cols-2">
                <div className='justify-items-center'>
                  <h3 className="text-sm font-medium text-gray-500 uppercase mb-2">Last Seen Location</h3>
                  <p className="text-gray-800 text-center px-2">{post.lastSeen?.address}</p>
                </div>
                <div className='justify-items-center'>
                  <h3 className="text-sm font-medium text-gray-500 uppercase mb-2">Area</h3>
                  <p className="text-gray-800 text-center px-2">{post.area?.address}</p>
                </div>
              </div>
            </div>
          )}

          {post.type === 'found' && (
            <div className="cursor-pointer pb-6 pt-3" onClick={handleContentClick}>
              <div className="grid grid-cols-2">
                <div className='justify-items-center'>
                  <h3 className="text-sm font-medium text-gray-500 uppercase mb-2">Found Location</h3>
                  <p className="text-gray-800 text-center px-2">{post.lastSeen?.address}</p>
                </div>
                <div className='justify-items-center'>
                  <h3 className="text-sm font-medium text-gray-500 uppercase mb-2">Contact Information</h3>
                  <p className="text-gray-800 text-center px-2">{post.contactInfo}</p>
                </div>
              </div>
            </div>
          )}

          <div className="cursor-pointer">
            {post.medias?.length > 0 && (
              <MediaGallery media={post.medias} className="mb-3" />
            )}
          </div>
          
          <div className="flex items-center justify-between pt-3 border-t border-gray-100">
            <div className="flex space-x-3 w-full">
              <button className="flex items-center text-sm text-gray-600 hover:text-primary-600 transition-colors" 
                onClick={handleInteract}
              >
                { userInteracted ? 
                  <Heart fill="red" size={18} className="mr-1 text-red-500" /> :
                  <Heart size={18} className="mr-1" /> 
                }
                <span>{ userInteracted ? interactionCount + 1 : interactionCount }</span>
              </button>
              <button 
                className="flex items-center text-sm text-gray-600 hover:text-primary-600 transition-colors"
                onClick={handleContentClick}
              >
                <MessageCircle size={18} className="mr-1" />
                <span>{post.commentNum}</span>
              </button>
              <div className="grow"/>
              { post.type !== 'blog' && 
              <button 
                className="flex items-center text-sm text-gray-600 hover:text-gray-800 transition-colors"
                onClick={handleContentClick}
              >
                <span className='mr-2'>{post.participants?.length ?? 'no '} helpers</span>
                <Users size={18} className="mr-2" />
              </button> }
            </div>
          </div>
        </div>
      </Card>

      {showEditModal && (
        <EditPostModal
          post={post}
          onClose={() => setShowEditModal(false)}
          onSave={handleSaveEdit}
        />
      )}
    </>
  );
};

export default PostCard;