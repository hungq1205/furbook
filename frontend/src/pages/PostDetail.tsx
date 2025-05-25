import React, { useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';
import { ChevronLeft, Heart, MessageCircle, Users, AlertTriangle, Copy } from 'lucide-react';
import { motion } from 'framer-motion';
import Button from '../components/common/Button';
import Avatar from '../components/common/Avatar';
import Card from '../components/common/Card';
import MediaGallery from '../components/feed/MediaGallery';
// import { posts } from '../data/mockData';
import { formatDistanceToNow } from '../utils/dateUtils';
import { handleError } from '../utils/errors';
import { postService } from '../services/postService';
import { Post, Comment } from '../types/post';
import { useAuth } from '../services/authService';

const PostDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const authService = useAuth();

  const [post, setPost] = React.useState<Post | null>(null);
  const [comments, setComments] = React.useState<Comment[]>([]); 
  const [comment, setComment] = React.useState<string>('');

  const [userInteracted, setUserInteracted] = React.useState<Boolean>(false);
  const [interactionCount, setInteractionCount] = React.useState<number>(0);

  const [userHelped, setUserHelped] = React.useState<Boolean>(false);
  const [participantCount, setParticipantCount] = React.useState<number>(0);

  const authUsername = authService.currentUser!.username;

  const fetchComments = () => {
    if (!id) return;
    postService.getComments(id!)
      .then(comments => setComments(comments.reverse()))
      .catch(error => handleError(error, 'Failed to fetch comments', authService.logout));
  };

  const handleInteract = () => {
    if (!id) return;
    if (userInteracted)
      postService.deleteInteraction(id)
        .then(() => userInteracted && setUserInteracted(false))
        .catch(error => handleError(error, 'Failed to unlike post', authService.logout));
    else
      postService.upsertInteraction(id)
        .then(() => !userInteracted && setUserInteracted(true))
        .catch(error => handleError(error, 'Failed to like post', authService.logout));
    setUserInteracted(!userInteracted);
  };

  const handleParticipate = () => {
    if (!id) return;
    if (post?.type === 'blog') return;
    if (userHelped)
      postService.unparticipate(id)
        .then(() => userHelped && setUserHelped(false))
        .catch(error => handleError(error, 'Failed to remove participation', authService.logout));
    else
      postService.participate(id)
        .then(() => !userHelped && setUserHelped(true))
        .catch(error => handleError(error, 'Failed to participate in post', authService.logout));
    setUserHelped(!userHelped);
  };

  const handleComment = () => {
    if (!id || !comment.trim()) return;
    postService.addComment(id, comment.trim())
      .then(() => fetchComments())
      .catch(error => handleError(error, 'Failed to post comment', authService.logout));
    setComment('');
  };

  useEffect(() => {
    postService.getById(id!)
      .then(post => {
        const isUserInteracted = post.interactions.some(interaction => interaction.username === authUsername);
        setPost(post);
        setUserInteracted(isUserInteracted);
        setInteractionCount(isUserInteracted ? post.interactions.length - 1 : post.interactions.length);

        if (post.type !== 'blog' && post.participants) {
          const isUserHelping = post.participants.some(participant => participant === authUsername);
          setUserHelped(isUserHelping);
          setParticipantCount(isUserHelping ? post.participants.length - 1 : post.participants.length);
        }
      })
      .catch(error => handleError(error, 'Failed to fetch post', authService.logout));
    fetchComments();
  }, [id]);

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
          
          {post.type !== 'blog' && (
            <>
              <div className="mb-4 inline-block px-3 py-1 rounded-full text-sm font-medium bg-error-100 text-error-700">
                {post.type === 'lost' ? 'Missing' : 'Found'}
              </div>
            </>
          )}
          
          <p className="text-gray-800 text-lg mb-6">{post.content}</p>
          
          {post.type !== 'blog' && (
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
                    <p className="text-gray-800">{userHelped ? participantCount + 1 : participantCount} people helping</p>
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
            <div className="flex space-x-6 mb-2 md:mb-0">
              <button className="flex items-center text-gray-600 hover:text-red-500" 
                onClick={handleInteract}
              >
                { userInteracted ? 
                  <Heart fill="red" size={20} className="mr-2 text-red-500" /> :
                  <Heart size={20} className="mr-2" /> 
                }
                <span>{userInteracted ? interactionCount + 1 : interactionCount} Likes</span>
              </button>
              <div className="flex items-center text-gray-600 cursor-default">
                <MessageCircle size={20} className="mr-2" />
                <span>{comments.length} Comments</span>
              </div>
            </div>
            {post.type !== 'blog' && (
              <div className="flex space-x-3">
                <Button
                  variant={post.isResolved ? 'ghost' : 'secondary'}
                  icon={<Users size={18} />}
                  disabled={post.isResolved}
                  onClick={handleParticipate}
                >
                  {post.isResolved ? 'Resolved' : userHelped ? 'Helping' : 'Help Find'}
                </Button>
                <Button variant="outline" ring={false} icon={<Copy size={18} />}>
                  Copy Link
                </Button>
              </div>
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
                  onChange={e => setComment(e.target.value)}
                  value={comment}
                ></textarea>
                <div className="flex justify-end mt-2">
                  <Button variant="primary" size="sm" onClick={handleComment}>Post Comment</Button>
                </div>
              </div>
            </div>
            
            {comments.length > 0 ? (
                <div className="space-y-4">
                {comments.map((comment, index) => (
                  <div key={index} className="flex space-x-3">
                    <Link to={`/profile/${comment.username}`}>
                      <Avatar 
                        src={comment.avatar} 
                        alt={comment.displayName} 
                        size="md" 
                      />
                    </Link>
                    <div>
                      <div className="bg-gray-100 p-3 rounded-lg justify-items-start">
                        <Link to={`/profile/${comment.username}`} className="mb-1 hover:underline">
                          <div className="font-medium text-gray-900">{comment.displayName}</div>
                        </Link>
                        <p className="text-gray-800">{comment.content}</p>
                      </div>
                      <div className="flex items-center mt-1 text-xs text-gray-500 ml-2">
                        <span>{formatDistanceToNow(new Date(comment.createdAt))}</span>
                      </div>
                    </div>
                  </div>
                ))}
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