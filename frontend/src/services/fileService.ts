import { HttpError } from './util';

const FILE_SERVICE_URL = `https://api.imgbb.com/1/upload`;

export const fileService = {
  async upload(file: File): Promise<string> {
    const formData = new FormData();
    formData.append('image', file);
    formData.append('key', import.meta.env.VITE_IMGBB_API_KEY || 'e32a70744a858799efb1317aab023f06');

    const response = await fetch(FILE_SERVICE_URL, {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) throw new HttpError(response.status, await response.json());
    const data = await response.json();
    return data.data.url;
  },
};