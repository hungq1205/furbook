import React, { useState } from 'react';
import { Image, Video, Send } from 'lucide-react';
import Avatar from '../common/Avatar';
import Button from '../common/Button';
import Card from '../common/Card';
// import { currentUser } from '../../data/mockData';
import { motion } from 'framer-motion';
import { useAuth } from '../../services/authService';
import { Media, postService, BlogPostPayload } from '../../services/postService';
import { fileService } from '../../services/fileService';

const CreatePostInput: React.FC = () => {
  const authService = useAuth();

  const [content, setContent] = useState('');
  const [expanded, setExpanded] = useState(false);
  const [uploadMedias, setUploadMedias]  = useState<File[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  const handleFocus = () => {
    setExpanded(true);
  };

  const handleCancel = () => {
    setContent('');
    setExpanded(false);
  };

  const handleSubmit = async () => {
    if (isLoading) return;
    setIsLoading(true);
    const medias: Media[] = await Promise.all(
      uploadMedias.map(async (file) => ({
        type: file.type.startsWith('image/') ? 'image' : 'video',
        url: await fileService.upload(file),
      } as Media))
    );
    await postService.createBlogPost({ content, medias } as BlogPostPayload);

    setContent('');
    setUploadMedias([]);
    setExpanded(false);
    setIsLoading(false);
  };

  const handleImageUpload = () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = 'image/*';
    input.onchange = (e: any) => {
      const file = e.target.files[0];
      file && setUploadMedias((prev) => [...prev, file]);
    };
    input.click();
  }

  const handleVideoUpload = () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = 'video/*';
    input.onchange = (e: any) => {
      const file = e.target.files[0];
      file && setUploadMedias((prev) => [...prev, file]);
    };
    input.click();
  }

  return (
    <Card className="mb-6 p-4 bg-white">
      <div className="flex items-start space-x-3">
        <Avatar src={authService.currentUser!.avatar} alt={authService.currentUser!.displayName} size="md" />
        
        <div className="flex-1">
          <div 
            className="border rounded-lg p-3 w-full bg-gray-50 hover:bg-white focus-within:bg-white transition-colors"
          >
            <textarea
              placeholder="Share what your pet is up to..."
              value={content}
              disabled={isLoading}
              onChange={(e) => setContent(e.target.value)}
              onClick={handleFocus}
              className="w-full bg-transparent border-none resize-none focus:outline-none min-h-[15px]"
              rows={expanded ? 2 : 1}
            />
            
            {expanded && (
              <motion.div 
                className="flex flex-col space-y-2 my-1"
                initial={{ opacity: 0, height: 0 }}
                animate={{ opacity: 1, height: 'auto' }}
                transition={{ duration: 0.3 }}
              >
                <div className="flex space-x-2">
                  <button 
                    disabled={isLoading}
                    className="flex items-center text-gray-600 hover:text-primary-600 transition-colors text-sm"
                    onClick={handleImageUpload}
                  >
                    <Image size={18} className="mr-1" />
                    <span>Photo</span>
                  </button>
                  <button 
                    disabled={isLoading}
                    className="flex items-center text-gray-600 hover:text-primary-600 transition-colors text-sm"
                    onClick={handleVideoUpload}
                  >
                    <Video size={18} className="mr-1" />
                    <span>Video</span>
                  </button>
                </div>
                <FileUploadPreview files={uploadMedias} setFiles={setUploadMedias} />
              </motion.div>
            )}
          </div>
          {expanded && (
            <motion.div 
              className="flex flex-col space-y-3 mt-3"
              initial={{ opacity: 0, height: 0 }}
              animate={{ opacity: 1, height: 'auto' }}
              transition={{ duration: 0.3 }}
            >
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
    </Card>
  );
};

function FileUploadPreview({ files, setFiles }: { files: File[]; setFiles: (files: File[]) => void }) {
  return (
  <div className="flex flex-wrap gap-2">
    {files.map((media, index) => (
      <div key={index} className="relative w-20 h-20">
        <img
          src={URL.createObjectURL(media)}
          alt={`Uploaded media ${index + 1}`}
          className="w-full h-full object-cover rounded-lg"
        />
        <button
          className="absolute top-1 right-1 bg-red-500 text-white rounded-full p-1 text-xs flex items-center justify-center w-5 h-5"
          onClick={() => setFiles(files.filter((_, i) => i !== index)) }
        >
          ✕
        </button>
      </div>
    ))}
  </div>
  );
} 

export default CreatePostInput;