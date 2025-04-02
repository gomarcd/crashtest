export interface RequestConfig {
    method: string;
    url: string;
    headers: Record<string, string>;
    queryParams: Record<string, string>;
    body: string;
  }
  
  export interface APIResponse {
    statusCode: number;
    headers: Record<string, string>;
    body: any;
    timeMs: number;
    error?: string;
  }
  
  // Header types for UI
  export interface Header {
    key: string;
    value: string;
    enabled: boolean;
  }
  
  // Query param types for UI
  export interface QueryParam {
    key: string;
    value: string;
    enabled: boolean;
  }
  
  // Collection and Request types
  export interface Collection {
    id: string;
    name: string;
    expanded: boolean;
    requests: Request[];
  }
  
  export interface Request {
    id: number | string;
    name: string;
    method: string;
    url: string;
    body?: string;
    headers?: Record<string, string>;
    params?: Record<string, string>;
    auth?: {
      type: 'none' | 'bearer' | 'basic' | 'apikey';
      config: any;
    };
    lastStatus?: number;
  }
  
  export interface RequestMethod {
    name: string;
    color: string;
  }
  
  export const REQUEST_METHODS: RequestMethod[] = [
    { name: 'GET', color: 'bg-green-500' },
    { name: 'POST', color: 'bg-yellow-500' },
    { name: 'PUT', color: 'bg-blue-500' },
    { name: 'DELETE', color: 'bg-red-500' },
    { name: 'PATCH', color: 'bg-purple-500' },
    { name: 'HEAD', color: 'bg-gray-500' },
    { name: 'OPTIONS', color: 'bg-indigo-500' }
  ];