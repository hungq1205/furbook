import React, { useState } from 'react';
import { Image, Video, Send } from 'lucide-react';
import Avatar from '../common/Avatar';
import Button from '../common/Button';
import Card from '../common/Card';
// import { currentUser } from '../../data/mockData';
import { motion } from 'framer-motion';
import { authService } from '../../services/authService';

const CreatePostInput: React.FC = () => {
  const [content, setContent] = useState('');
  const [expanded, setExpanded] = useState(false);

  const handleFocus = () => {
    setExpanded(true);
  };

  const handleCancel = () => {
    setContent('');
    setExpanded(false);
  };

  const handleSubmit = () => {
    // In a real app, we would submit the post to an API
    console.log('Posting:', content);
    setContent('');
    setExpanded(false);
  };

  return (
    <Card className="mb-6 p-4 bg-white">
      <div className="flex items-start space-x-3">
        <Avatar src={authService.getCurrentUser().avatar} alt={authService.getCurrentUser().displayName} size="md" />
        
        <div className="flex-1">
          <div 
            className="border rounded-lg p-3 w-full bg-gray-50 hover:bg-white focus-within:bg-white transition-colors"
            onClick={handleFocus}
          >
            <textarea
              placeholder="Share what your pet is up to..."
              value={content}
              onChange={(e) => setContent(e.target.value)}
              onFocus={handleFocus}
              className="w-full bg-transparent border-none resize-none focus:outline-none min-h-[40px]"
              rows={expanded ? 3 : 1}
            />
            
            {expanded && (
              <motion.div 
                className="flex flex-col space-y-3 mt-3"
                initial={{ opacity: 0, height: 0 }}
                animate={{ opacity: 1, height: 'auto' }}
                transition={{ duration: 0.3 }}
              >
                <div className="flex space-x-2">
                  <button className="flex items-center text-gray-600 hover:text-primary-600 transition-colors text-sm">
                    <Image size={18} className="mr-1" />
                    <span>Photo</span>
                  </button>
                  <button className="flex items-center text-gray-600 hover:text-primary-600 transition-colors text-sm">
                    <Video size={18} className="mr-1" />
                    <span>Video</span>
                  </button>
                </div>
                
                <div className="flex justify-end space-x-2">
                  <Button variant="ghost" onClick={handleCancel}>
                    Cancel
                  </Button>
                  <Button 
                    variant="primary" 
                    onClick={handleSubmit}
                    disabled={!content.trim()}
                    icon={<Send size={16} />}
                  >
                    Post
                  </Button>
                </div>
              </motion.div>
            )}
          </div>
        </div>
      </div>
    </Card>
  );
};

export default CreatePostInput;