import React, { useState } from 'react';
import { Heart, MessageCircle, Share2, MoreHorizontal, Edit, Trash2 } from 'lucide-react';
import { Link, useNavigate } from 'react-router-dom';
import { Post } from '../../types/post';
import Avatar from '../common/Avatar';
import Card from '../common/Card';
import IconButton from '../common/IconButton';
import MediaGallery from './MediaGallery';
import EditPostModal from '../post/EditPostModal';
import { formatDistanceToNow } from '../../utils/dateUtils';
// import { currentUser } from '../../data/mockData';
import { postService } from '../../services/postService';
import { authService } from '../../services/authService';

interface PostCardProps {
  post: Post;
  onPostUpdated?: (post: Post) => void;
  onPostDeleted?: (postId: string) => void;
}

const PostCard: React.FC<PostCardProps> = ({ post, onPostUpdated, onPostDeleted }) => {
  const navigate = useNavigate();
  const [showOptions, setShowOptions] = useState(false);
  const [showEditModal, setShowEditModal] = useState(false);
  const isOwnPost = post.username === authService.getCurrentUser().username;

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
        onPostDeleted?.(post.id);
      } catch (error) {
        console.error('Failed to delete post:', error);
      }
    }
    setShowOptions(false);
  };

  const handleSaveEdit = async (updatedPost: Post) => {
    onPostUpdated?.(updatedPost);
    setShowEditModal(false);
  };

  const likeCount = post.interactions.filter(i => i.type === 'like').length;
  const shareCount = post.interactions.filter(i => i.type === 'share').length;

  return (
    <>
      <Card className="mb-4">
        <div className="p-4">
          <div className="flex items-center justify-between mb-3">
            <div className="flex items-center space-x-3">
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
          
          <div onClick={handleContentClick} className="cursor-pointer">
            <p className="text-gray-800 mb-3">{post.content}</p>
            
            {post.medias?.length > 0 && (
              <MediaGallery media={post.medias} className="mb-3" />
            )}
          </div>
          
          <div className="flex items-center justify-between pt-3 border-t border-gray-100">
            <div className="flex space-x-3">
              <button className="flex items-center text-sm text-gray-600 hover:text-primary-600 transition-colors">
                <Heart size={18} className="mr-1" />
                <span>{likeCount}</span>
              </button>
              <button 
                className="flex items-center text-sm text-gray-600 hover:text-primary-600 transition-colors"
                onClick={handleContentClick}
              >
                <MessageCircle size={18} className="mr-1" />
                <span>{post.commentNum}</span>
              </button>
              <button className="flex items-center text-sm text-gray-600 hover:text-primary-600 transition-colors">
                <Share2 size={18} className="mr-1" />
                <span>{shareCount}</span>
              </button>
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