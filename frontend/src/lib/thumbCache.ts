import { GetThumbnail } from '../../wailsjs/go/main/App';

const cache = new Map<string, string>();
const inflight = new Map<string, Promise<string>>();

export async function getThumbUrl(assetId: string): Promise<string> {
  if (cache.has(assetId)) return cache.get(assetId)!;
  if (inflight.has(assetId)) return inflight.get(assetId)!;

  const promise = (async () => {
    const b64 = await GetThumbnail(assetId) as unknown as string;
    if (!b64) return '';
    const binary = atob(b64);
    const buf = new Uint8Array(binary.length);
    for (let i = 0; i < binary.length; i++) buf[i] = binary.charCodeAt(i);
    const url = URL.createObjectURL(new Blob([buf]));
    cache.set(assetId, url);
    inflight.delete(assetId);
    return url;
  })();

  inflight.set(assetId, promise);
  return promise;
}

export function getCachedUrl(assetId: string): string | undefined {
  return cache.get(assetId);
}
