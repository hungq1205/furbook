import React, { useEffect, useState } from 'react';
import { motion } from 'framer-motion';
import { Upload, X, MapPin } from 'lucide-react';
import { Location, Media } from '../types/post';
import Button from '../components/common/Button';
import Card from '../components/common/Card';
import { fileService } from '../services/fileService';
import LocationPicker from '../components/map/LocationPicker';
import { LostPostPayload, postService } from '../services/postService';

const CreateLostPet: React.FC = () => {
  const [userLocation, setUserLocation] = useState<{ lat: number; lng: number } | undefined>();
  const [type, setType] = useState<'lost' | 'found'>('lost');
  const [previewUrls, setPreviewUrls] = useState<string[]>([]);
  const [showLocationPicker, setShowLocationPicker] = useState<'area' | 'lastSeen' | null>(null);
  const [area, setArea] = useState<Location | null>(null);
  const [lastSeen, setLastSeen] = useState<Location | null>(null);
  const [isLoading, setIsLoading] = useState(false);

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

  const handleLocationSelect = (loc: Location) => {
    if (showLocationPicker === 'area')
      setArea(loc);
    else
      setLastSeen(loc);
    setShowLocationPicker(null);
  };

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

  const toMedias = async (files: File[]): Promise<Media[]> => {
    try {
      return await Promise.all(
        files.map(async file => {
          const url = await fileService.upload(file);
          return {
            type: file.type.startsWith('image/') ? 'image' : 'video',
            url: url,
          } as Media;
        })
      );
    } catch (error) {
      console.error('Failed to upload files:', error);
    }
    return [];
  }

  const submitForm = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);

    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);

    const rawLostAt = formData.get('lostAt') as string;
    const lostAt = new Date(rawLostAt).toISOString();

    const data = {
      type: type,
      content: formData.get('description') as string,
      lastSeen: lastSeen,
      area: area,
      contactInfo: formData.get('contact') as string,
      lostAt: lostAt,
      medias: await toMedias(formData.getAll('medias') as File[]),
    } as LostPostPayload;

    try {
      await postService.createLostPost(data);
    } catch (err) {
      console.error('Failed to post lost found post:', err);
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <motion.div
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
    >
      <h1 className="text-2xl font-bold text-gray-900 mb-6">
        {type === 'lost' ? 'Report Lost Pet' : 'Report Found Pet'}
      </h1>
      
      <Card>
        <form className="p-6 space-y-6" onSubmit={submitForm}>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">Status</label>
            <div className="flex space-x-4">
              <button
                type="button"
                className={`px-4 py-2 rounded-md border focus:outline-none ${
                  type === 'lost'
                    ? 'bg-error-50 border-error-200 text-error-700'
                    : 'border-gray-300 text-gray-700 hover:bg-gray-50'
                }`}
                onClick={() => setType('lost')}
              >
                Lost Pet
              </button>
              <button
                type="button"
                className={`px-4 py-2 rounded-md border focus:outline-none ${
                  type === 'found'
                    ? 'bg-orange-50 border-orange-200 text-orange-700'
                    : 'border-gray-300 text-gray-700 hover:bg-gray-50'
                }`}
                onClick={() => setType('found')}
              >
                Found Pet
              </button>
            </div>
          </div>
          
          <div>
            <label htmlFor="description" className="block text-sm font-medium text-gray-700 mb-2">Description</label>
            <textarea
              name="description"
              rows={4}
              placeholder="Describe your pet in detail including color, breed, size, distinctive features, etc."
              className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
            ></textarea>
          </div>
          
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="flex flex-col md:col-span-2">
              <label htmlFor="lastSeen" className="block text-sm font-medium text-gray-700 mb-2">
                { type === 'lost' ? 'Last Seen Location' : 'Found Location' }
              </label>
              <div className="relative grow">
                <button
                  type="button"
                  onClick={() => setShowLocationPicker('lastSeen')}
                  className="w-full flex items-center px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500 bg-white"
                >
                  <MapPin size={16} className="text-gray-400 mr-2" />
                  {lastSeen ? lastSeen.address : 'Select on map'}
                </button>
              </div>
            </div>

            { type === 'lost' && 
            <div className="flex flex-col md:col-span-2">
              <label htmlFor="area" className="block text-sm font-medium mb-2">Area</label>
              <div className="relative grow">
                <button
                  type="button"
                  onClick={() => setShowLocationPicker('area')}
                  className="w-full flex items-center px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500 bg-white"
                >
                  <MapPin size={16} className="text-gray-400 mr-2" />
                  {area ? area.address : 'Select on map'}
                </button>
              </div>
            </div>
            }

            <div>
              <label htmlFor="contact" className="block text-sm font-medium text-gray-700 mb-2">Contact Information</label>
              <input
                type="text"
                name="contact"
                placeholder="Phone number, email, etc."
                className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
            
            <div>
              <label htmlFor="lostAt" className="block text-sm font-medium text-gray-700 mb-2">
                {type === 'lost' ? 'Date Lost' : 'Date Found'}
              </label>
              <input
                type="date"
                name="lostAt"
                className="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
              />
            </div>
          </div>
          
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">Photos/Videos</label>
            {previewUrls.length > 0 && (
              <div className="mt-4 grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-2 mb-6">
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
            
            <div className="border-2 border-dashed border-gray-300 rounded-md p-6 flex flex-col items-center">
              <div className="mb-3 text-gray-500">
                <Upload size={24} />
              </div>
              <p className="text-sm text-gray-500 mb-2">Drag and drop files here or click to upload</p>
              <p className="text-xs text-gray-400 mb-3">Images help others identify your pet</p>
              <input
                type="file"
                name="medias"
                multiple
                accept="image/*,video/*"
                onChange={handleFileChange}
                className="hidden"
              />
              <label htmlFor="medias">
                <Button variant="outline" size="sm" type="button" onClick={() => document.getElementsByName('medias')[0]?.click()}>
                  Select Files
                </Button>
              </label>
            </div>
          </div>
          
          <div className="flex justify-end space-x-3 pt-4 border-t border-gray-200">
            <Button variant="ghost" type="button" onClick={() => window.history.back()}>
              Cancel
            </Button>
            <Button variant="primary" type="submit">Report Pet</Button>
          </div>
        </form>
      </Card>

      {showLocationPicker && (
        <LocationPicker
          onSelected={handleLocationSelect}
          onClose={() => setShowLocationPicker(null)}
          userLocation={userLocation}
        />
      )}

      {isLoading && (
        <div className="fixed inset-0 z-50 bg-black/10 flex items-center justify-center">
          <div className="animate-spin rounded-full h-12 w-12 border-t-4 border-primary-500 border-opacity-50"></div>
        </div>
      )}
    </motion.div>
  );
};

export default CreateLostPet;