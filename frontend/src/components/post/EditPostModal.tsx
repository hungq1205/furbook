import React, { useState } from 'react';
import { X, Upload, Trash2 } from 'lucide-react';
import { Post, Media } from '../../types/post';
import Button from '../common/Button';

interface EditPostModalProps {
  post: Post;
  onClose: () => void;
  onSave: (updatedPost: Post) => void;
}

const EditPostModal: React.FC<EditPostModalProps> = ({ post, onClose, onSave }) => {
  const [content, setContent] = useState(post.content);
  const [medias, setMedias] = useState<Media[]>(post.medias || []);
  const [newMedia, setNewMedia] = useState<File[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const filesArray = Array.from(e.target.files);
      setNewMedia([...newMedia, ...filesArray]);

      const newMediaPreviews: Media[] = filesArray.map(file => ({
        type: file.type.startsWith('image/') ? 'image' : 'video',
        url: URL.createObjectURL(file)
      }));

      setMedias([...medias, ...newMediaPreviews]);
    }
  };

  const removeMedia = (index: number) => {
    setMedias(medias.filter((_, i) => i !== index));
    if (index >= post.medias.length) {
      const newMediaIndex = index - post.medias.length;
      setNewMedia(newMedia.filter((_, i) => i !== newMediaIndex));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      // First, upload any new media files
      const mediaUrls = await Promise.all(
        newMedia.map(async (file) => {
          const formData = new FormData();
          formData.append('file', file);
          const response = await fetch('/api/upload', {
            method: 'POST',
            body: formData
          });
          if (!response.ok) throw new Error('Failed to upload media');
          return response.json();
        })
      );

      // Update the post with new content and media
      // const updatedPost = await postApi.update(post.id, {
      //   content,
      //   medias: [
      //     ...post.medias,
      //     ...mediaUrls.map(url => ({
      //       id: Math.random().toString(),
      //       type: url.includes('.mp4') ? 'video' : 'image',
      //       url
      //     }))
      //   ]
      // });

      // onSave(updatedPost);
      onClose();
    } catch (error) {
      console.error('Failed to update post:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg w-full max-w-2xl mx-4">
        <div className="flex items-center justify-between p-4 border-b">
          <h2 className="text-xl font-semibold">Edit Post</h2>
          <button onClick={onClose} className="text-gray-500 hover:text-gray-700">
            <X size={24} />
          </button>
        </div>

        <form onSubmit={handleSubmit} className="p-4">
          <textarea
            value={content}
            onChange={(e) => setContent(e.target.value)}
            className="w-full p-3 border rounded-lg resize-none focus:ring-2 focus:ring-primary-500 focus:border-transparent"
            rows={4}
            placeholder="What's on your mind?"
          />

          {medias.length > 0 && (
            <div className="mt-4 grid grid-cols-2 sm:grid-cols-3 gap-2">
              {medias.map((item, index) => (
                <div key={item.url} className="relative aspect-square">
                  {item.type === 'image' ? (
                    <img
                      src={item.url}
                      alt="Post media"
                      className="w-full h-full object-cover rounded-lg"
                    />
                  ) : (
                    <video
                      src={item.url}
                      className="w-full h-full object-cover rounded-lg"
                    />
                  )}
                  <button
                    type="button"
                    onClick={() => removeMedia(index)}
                    className="absolute top-1 right-1 p-1 bg-white rounded-full shadow-sm hover:bg-gray-100"
                  >
                    <Trash2 size={16} className="text-red-500" />
                  </button>
                </div>
              ))}
            </div>
          )}

          <div className="mt-4">
            <input
              type="file"
              id="media"
              multiple
              accept="image/*,video/*"
              onChange={handleFileChange}
              className="hidden"
            />
            <label
              htmlFor="media"
              className="flex items-center justify-center p-4 border-2 border-dashed border-gray-300 rounded-lg cursor-pointer hover:border-primary-500"
            >
              <Upload size={24} className="mr-2 text-gray-500" />
              <span className="text-gray-600">Add Photos/Videos</span>
            </label>
          </div>

          <div className="flex justify-end space-x-3 mt-4 pt-4 border-t">
            <Button variant="ghost" onClick={onClose} disabled={isLoading}>
              Cancel
            </Button>
            <Button
              variant="primary"
              type="submit"
              disabled={isLoading || !content.trim()}
            >
              {isLoading ? 'Saving...' : 'Save Changes'}
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default EditPostModal;