import { HttpError } from './util';

const UPLOADCARE_URL = 'https://upload.uploadcare.com/base/';
const PUBLIC_KEY = import.meta.env.VITE_UPLOADCARE_PUBLIC_KEY || '35f10945513f36b879bc';

export const fileService = {
  async upload(file: File): Promise<string> {
    const formData = new FormData();
    formData.append('UPLOADCARE_PUB_KEY', PUBLIC_KEY);
    formData.append('UPLOADCARE_STORE', '1');
    formData.append('file', file);

    const response = await fetch(UPLOADCARE_URL, {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) throw new HttpError(response.status, await response.json());
    const data = await response.json();
    return `https://ucarecdn.com/${data.file}/`;
  },
};
