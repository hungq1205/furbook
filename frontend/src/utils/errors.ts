import { HttpError } from "../services/util";

export const handleError = (error: any, message: string, logout: () => void) => {
    if (error instanceof HttpError) {
        console.error(`Status ${error.status}, ${message}: ${error.message}`);
        if (error.status === 401) {
            logout();
            window.location.href = '/login';
        }
    } else if (error instanceof Error) {
        console.error(`${message}: ${error.message}`);
    } else {
        console.error(`Unknown error of ${message}: ${error}`);
    }
}