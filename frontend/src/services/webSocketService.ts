// WebSocketService.ts
type MessageType = 'auth' | 'chat' | 'notification';

export type Message = {
  type: MessageType;
  payload: any;
};

type AuthPayload = {
  token: string;
};

type AuthResponsePayload = {
  status: string;
};

export type ChatPayload = {
  messageId: number;
  username: string;
  groupId: number;
  content: string;
  createdAt: string;
};

export type NotificationPayload = {
  title: string;
  body: string;
};

type OnMessageCallback<T = any> = (payload: T) => void;

class WebSocketService {
  private socket: WebSocket | null = null;
  private messageListeners: Map<MessageType, OnMessageCallback[]> = new Map();
  private isAuthenticated: boolean = false;
  private authPromise: Promise<void> | null = null;

  async connect(token: string): Promise<void> {
    const wsUrl = `ws://localhost:3000/ws`;
    this.socket = new WebSocket(wsUrl);
    
    this.authPromise = new Promise((resolve, reject) => {
      if (!this.socket) {
        reject(new Error('Failed to create WebSocket connection'));
        return;
      }

      this.socket.onopen = () => {
        console.log("[WebSocket] Connected, sending auth...");
        
        const authMessage: Message = {
          type: 'auth',
          payload: { token } as AuthPayload
        };
        this.socket!.send(JSON.stringify(authMessage));
      };

      this.socket.onmessage = (event: MessageEvent) => {
        try {
          const msg: Message = JSON.parse(event.data);

          if (msg.type === 'auth' && !this.isAuthenticated) {
            const authResponse = msg.payload as AuthResponsePayload;
            if (authResponse.status === 'success') {
              console.log("[WebSocket] Authenticated successfully");
              this.isAuthenticated = true;
              resolve();
            } else {
              console.error("[WebSocket] Authentication failed");
              reject(new Error('Authentication failed'));
            }
            return;
          }

          if (this.isAuthenticated)
            this.notifyListeners(msg.type, msg.payload);
        } catch (err) {
          console.error("Failed to parse message:", err);
        }
      };

      this.socket.onerror = (err) => {
        console.error("[WebSocket] Error:", err);
        reject(err);
      };

      this.socket.onclose = () => {
        console.log("[WebSocket] Disconnected");
        this.socket = null;
        this.isAuthenticated = false;
        this.authPromise = null;
      };
    });

    await this.authPromise;
  }

  async send(message: Message): Promise<void> {
    if (this.authPromise && !this.isAuthenticated) {
      await this.authPromise;
    }

    if (!this.isAuthenticated) {
      throw new Error('WebSocket is not authenticated');
    }

    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify(message));
    } else {
      throw new Error('WebSocket is not open');
    }
  }

  subscribe<T>(type: MessageType, callback: (payload: T) => void): void {
    const listeners = this.messageListeners.get(type) || [];
    listeners.push(callback);
    this.messageListeners.set(type, listeners);
  }

  unsubscribe(type: MessageType, callback: OnMessageCallback): void {
    const callbacks = this.messageListeners.get(type);
    if (!callbacks) return;
    this.messageListeners.set(type, callbacks.filter(cb => cb !== callback));
  }

  disconnect(): void {
    this.socket?.close();
    this.socket = null;
    this.isAuthenticated = false;
    this.authPromise = null;
    this.messageListeners = new Map();
  }

  private notifyListeners(type: MessageType, payload: any): void {
    const listeners = this.messageListeners.get(type);
    if (!listeners) return;
    listeners.forEach(callback => callback(payload));
  }
  
  get connected(): boolean {
    return this.socket?.readyState === WebSocket.OPEN && this.isAuthenticated;
  }
}

const wsService = new WebSocketService();
export default wsService;