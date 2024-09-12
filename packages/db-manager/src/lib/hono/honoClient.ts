import type { AppType } from './hono.server';
import { hc } from 'hono/client';

export const honoClient = hc<AppType>('http://localhost:5100/');
