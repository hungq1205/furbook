import { HttpError } from './util';

// TODO: replace with actual url
const CLOUDINARY_URL = 'https://api.cloudinary.com/v1_1/<your-cloud-name>/upload';
const CLOUDINARY_UPLOAD_PRESET = '<your-upload-preset>';

export const fileService = {
  async upload(file: File): Promise<string> {
    const formData = new FormData();
    formData.append('file', file);
    formData.append('upload_preset', CLOUDINARY_UPLOAD_PRESET);

    const response = await fetch(CLOUDINARY_URL, {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) throw new HttpError(response.status, await response.json());
    const data = await response.json();
    return data.secure_url;
  },
};