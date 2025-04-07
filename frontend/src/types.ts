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
  body: string | object | null;
  timeMs: number;
  error?: string;
}

export interface Header {
  key: string;
  value: string;
  enabled: boolean;
}

export interface QueryParam {
  key: string;
  value: string;
  enabled: boolean;
}

export interface Collection {
  id: string;
  name: string;
  expanded: boolean;
  requests: Request[];
}

export type AuthConfig =
  | { type: 'none'; config?: never }
  | { type: 'bearer'; config: { token: string } }
  | { type: 'basic'; config: { username: string; password: string } }
  | { type: 'apikey'; config: { key: string; value: string } };

export interface Request {
  id: number | string;
  name: string;
  method: string;
  url: string;
  body?: string;
  headers?: Record<string, string>;
  params?: Record<string, string>;
  auth?: AuthConfig;
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
  { name: 'OPTIONS', color: 'bg-indigo-500' },
];