import type { RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = () => {
  return new Response(JSON.stringify({ OK: true }), {
    headers: { 'Content-Type': 'application/manifest+json' }
  });
