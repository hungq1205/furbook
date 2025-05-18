import { authService } from "../services/authService";
import { HttpError } from "../services/util";

export const handleError = (error: any, message: string) => {
    if (error instanceof HttpError) {
        console.error(`Status ${error.status}, ${message}: ${error.message}`);
        if (error.status === 401) {
            authService.logout();
            window.location.href = '/login';
        }
    } else if (error instanceof Error) {
        console.error(`${message}: ${error.message}`);
    } else {
        console.error(`Unknown error of ${message}: ${error}`);
    }
}