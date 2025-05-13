import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { Upload, X, MapPin } from 'lucide-react';
import Button from '../components/common/Button';
import Card from '../components/common/Card';

const CreateLostPet: React.FC = () => {
  const [status, setStatus] = useState<'lost' | 'found'>('lost');
  const [previewUrls, setPreviewUrls] = useState<string[]>([]);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const filesArray = Array.from(e.target.files);
      const newPreviewUrls = filesArray.map(file => URL.createObjectURL(file));
      setPreviewUrls([...previewUrls, ...newPreviewUrls]);
    }
  };

  const removePreview = (index: number) => {
    setPreviewUrls(previewUrls.filter((_, i) => i !== index));
  };

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
    >
      <h1 className="text-2xl font-bold text-gray-900 mb-6">
        {status === 'lost' ? 'Report Lost Pet' : 'Report Found Pet'}
      </h1>
      
      <Card>
        <form className="p-6 space-y-6">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">Status</label>
            <div className="flex space-x-4">
              <button
                type="button"
                className={`px-4 py-2 rounded-md border focus:outline-none ${
                  status === 'lost'
                    ? 'bg-error-50 border-error-200 text-error-700'
                    : 'border-gray-300 text-gray-700 hover:bg-gray-50'
                }`}
                onClick={() => setStatus('lost')}
              >
                Lost Pet
              </button>
              <button
                type="button"
                className={`px-4 py-2 rounded-md border focus:outline-none ${
                  status === 'found'
                    ? 'bg-success-50 border-success-200 text-success-700'
                    : 'border-gray-300 text-gray-700 hover:bg-gray-50'
                }`}
                onClick={() => setStatus('found')}
              >
                Found Pet
              </button>
            </div>
          </div>
          
          <div>
            <label htmlFor="title" className="block text-sm font-medium text-gray-700 mb-2">Title</label>
            <input
              type="text"
              id="title"
              placeholder={`${status === 'lost' ? 'Missing' : 'Found'} [Pet Type] - [Pet Name/Description]`}
              className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
            />
          </div>
          
          <div>
            <label htmlFor="description" className="block text-sm font-medium text-gray-700 mb-2">Description</label>
            <textarea
              id="description"
              rows={4}
              placeholder="Describe your pet in detail including color, breed, size, distinctive features, etc."
              className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
            ></textarea>
          </div>
          
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label htmlFor="location" className="block text-sm font-medium text-gray-700 mb-2">
                Last Seen Location
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <MapPin size={16} className="text-gray-400" />
                </div>
                <input
                  type="text"
                  id="location"
                  placeholder="Street name, landmark, etc."
                  className="w-full pl-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
                />
              </div>
            </div>
            
            <div>
              <label htmlFor="area" className="block text-sm font-medium text-gray-700 mb-2">Area</label>
              <input
                type="text"
                id="area"
                placeholder="Neighborhood, district, etc."
                className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
            
            <div>
              <label htmlFor="contact" className="block text-sm font-medium text-gray-700 mb-2">Contact Information</label>
              <input
                type="text"
                id="contact"
                placeholder="Phone number, email, etc."
                className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
            
            <div>
              <label htmlFor="date" className="block text-sm font-medium text-gray-700 mb-2">
                {status === 'lost' ? 'Date Lost' : 'Date Found'}
              </label>
              <input
                type="date"
                id="date"
                className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
          </div>
          
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">Photos/Videos</label>
            <div className="border-2 border-dashed border-gray-300 rounded-md p-6 flex flex-col items-center">
              <div className="mb-3 text-gray-500">
                <Upload size={24} />
              </div>
              <p className="text-sm text-gray-500 mb-2">Drag and drop files here or click to upload</p>
              <p className="text-xs text-gray-400 mb-3">Images help others identify your pet</p>
              <input
                type="file"
                id="media"
                multiple
                accept="image/*,video/*"
                className="hidden"
                onChange={handleFileChange}
              />
              <label htmlFor="media">
                <Button variant="outline" size="sm" type="button">
                  Select Files
                </Button>
              </label>
            </div>
            
            {previewUrls.length > 0 && (
              <div className="mt-4 grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-2">
                {previewUrls.map((url, index) => (
                  <div key={index} className="relative aspect-square">
                    <img src={url} alt="Preview" className="h-full w-full object-cover rounded-md" />
                    <button
                      type="button"
                      onClick={() => removePreview(index)}
                      className="absolute top-1 right-1 bg-white rounded-full p-1 shadow-sm hover:bg-gray-100"
                    >
                      <X size={16} />
                    </button>
                  </div>
                ))}
              </div>
            )}
          </div>
          
          <div className="flex justify-end space-x-3 pt-4 border-t border-gray-200">
            <Button variant="ghost" type="button">
              Cancel
            </Button>
            <Button variant="primary" type="button">
              {status === 'lost' ? 'Report Lost Pet' : 'Report Found Pet'}
            </Button>
          </div>
        </form>
      </Card>
    </motion.div>
  );
};

export default CreateLostPet;