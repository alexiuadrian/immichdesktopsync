<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { getThumbUrl } from '../lib/thumbCache';
  import { DownloadAsset, GetStreamPort } from '../../wailsjs/go/main/App';

  export interface LightboxAsset {
    id: string;
    type: string;
    originalPath: string;
    localDateTime?: string;
    fileCreatedAt?: string;
  }

  export let assets: LightboxAsset[] = [];
  export let index: number = 0;
  export let onClose: () => void = () => {};

  $: asset = assets[index];
  $: isVideo = asset?.type === 'VIDEO';
  $: filename = asset?.originalPath?.replace(/\\/g, '/').split('/').pop() ?? '';
  $: dateStr = formatDate(asset?.localDateTime ?? asset?.fileCreatedAt ?? '');

  let imgPromise: Promise<string> = Promise.resolve('');
  $: if (asset && !isVideo) imgPromise = getThumbUrl(asset.id);

  $: {
    if (index > 0) getThumbUrl(assets[index - 1].id);
    if (index < assets.length - 1) getThumbUrl(assets[index + 1].id);
  }

  let streamPort = 0;
  $: streamUrl = (isVideo && streamPort && asset)
    ? `http://127.0.0.1:${streamPort}/video/${asset.id}`
    : '';

  let scale = 1;
  let panX = 0;
  let panY = 0;
  let isDragging = false;
  let didDrag = false;
  let dragStartX = 0;
  let dragStartY = 0;
  let dragStartPanX = 0;
  let dragStartPanY = 0;

  $: zoomPct = Math.round(scale * 100);

  function resetZoom() { scale = 1; panX = 0; panY = 0; }

  function zoomBy(factor: number, cx = 0, cy = 0) {
    const newScale = Math.max(0.5, Math.min(10, scale * factor));
    panX = cx - (cx - panX) * (newScale / scale);
    panY = cy - (cy - panY) * (newScale / scale);
    scale = newScale;
  }

  function onWheel(e: WheelEvent) {
    e.preventDefault();
    const rect = (e.currentTarget as HTMLElement).getBoundingClientRect();
    const cx = e.clientX - rect.left - rect.width / 2;
    const cy = e.clientY - rect.top - rect.height / 2;
    const factor = e.deltaY < 0 ? 1.15 : 1 / 1.15;
    zoomBy(factor, cx, cy);
  }

  function onImgMouseDown(e: MouseEvent) {
    if (e.button !== 0) return;
    isDragging = true;
    didDrag = false;
    dragStartX = e.clientX;
    dragStartY = e.clientY;
    dragStartPanX = panX;
    dragStartPanY = panY;
    e.preventDefault();
  }

  function onImgMouseMove(e: MouseEvent) {
    if (!isDragging) return;
    const dx = e.clientX - dragStartX;
    const dy = e.clientY - dragStartY;
    if (Math.abs(dx) > 2 || Math.abs(dy) > 2) didDrag = true;
    panX = dragStartPanX + dx;
    panY = dragStartPanY + dy;
  }

  function onImgMouseUp() { isDragging = false; }

  function onImgDblClick() {
    if (scale !== 1) resetZoom(); else zoomBy(2);
  }

  function prev() { if (index > 0) { resetZoom(); index--; } }
  function next() { if (index < assets.length - 1) { resetZoom(); index++; } }

  function formatDate(iso: string): string {
    if (!iso) return '';
    try {
      return new Date(iso).toLocaleDateString(undefined, {
        year: 'numeric', month: 'long', day: 'numeric',
      });
    } catch { return ''; }
  }

  function onKey(e: KeyboardEvent) {
    if (e.key === 'Escape') onClose();
    else if (e.key === 'ArrowLeft') prev();
    else if (e.key === 'ArrowRight') next();
    else if (e.key === '+' || e.key === '=') zoomBy(1.25);
    else if (e.key === '-') zoomBy(1 / 1.25);
    else if (e.key === '0') resetZoom();
  }

  onMount(async () => {
    window.addEventListener('keydown', onKey);
    streamPort = await GetStreamPort();
  });
  onDestroy(() => window.removeEventListener('keydown', onKey));

  let downloading = false;
  let downloadMsg = '';
  $: { asset; downloadMsg = ''; }

  async function handleDownload() {
    if (!asset) return;
    downloading = true;
    downloadMsg = '';
    try {
      await DownloadAsset(asset.id, filename || `${asset.id}.bin`);
      downloadMsg = 'Saved';
      setTimeout(() => { downloadMsg = ''; }, 3000);
    } catch (e: any) {
      downloadMsg = String(e).replace(/^Error: /, '');
      setTimeout(() => { downloadMsg = ''; }, 5000);
    } finally {
      downloading = false;
    }
  }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
  class="fixed inset-0 z-50 flex items-center justify-center bg-black/90"
  on:click|self={onClose}
>
  <button
    class="absolute right-4 top-4 z-10 rounded-full bg-black/50 p-2 text-white transition hover:bg-white/20"
    on:click={onClose} title="Close (Esc)"
  >
    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
      <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
    </svg>
  </button>

  <div class="absolute right-14 top-4 z-10 flex items-center gap-2">
    {#if downloadMsg}
      <span class="text-xs {downloadMsg === 'Saved' ? 'text-green-400' : 'text-red-400'}">{downloadMsg}</span>
    {/if}
    <button
      class="rounded-full bg-black/50 p-2 text-white transition hover:bg-white/20 disabled:opacity-40"
      on:click={handleDownload} disabled={downloading} title="Download original"
    >
      {#if downloading}
        <svg class="h-5 w-5 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"/>
        </svg>
      {:else}
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5M16.5 12L12 16.5m0 0L7.5 12m4.5 4.5V3"/>
        </svg>
      {/if}
    </button>
  </div>

  {#if index > 0}
    <button
      class="absolute left-3 top-1/2 z-10 -translate-y-1/2 rounded-full bg-black/50 p-3 text-white transition hover:bg-white/20"
      on:click={prev} title="Previous (←)"
    >
      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
        <polyline points="15 18 9 12 15 6"/>
      </svg>
    </button>
  {/if}

  <div class="flex h-full w-full items-center justify-center px-16 py-14">
    {#if isVideo}
      <!-- svelte-ignore a11y-media-has-caption -->
      <video
        src={streamUrl}
        controls
        autoplay
        class="max-h-full max-w-full rounded shadow-2xl"
        on:click|stopPropagation
      ></video>

    {:else}
      {#await imgPromise}
        <div class="flex flex-col items-center gap-3 text-white/40">
          <svg class="h-10 w-10 animate-spin" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"/>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z"/>
          </svg>
          <span class="text-sm">Loading…</span>
        </div>
      {:then url}
        {#if url}
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          <div
            class="h-full w-full flex items-center justify-center overflow-hidden"
            style="cursor: {isDragging ? 'grabbing' : scale > 1 ? 'grab' : 'default'}"
            on:wheel|preventDefault={onWheel}
            on:mousedown={onImgMouseDown}
            on:mousemove={onImgMouseMove}
            on:mouseup={onImgMouseUp}
            on:mouseleave={onImgMouseUp}
            on:dblclick={onImgDblClick}
            on:click|stopPropagation={() => {}}
          >
            <img
              src={url}
              alt={filename}
              class="max-h-full max-w-full rounded object-contain shadow-2xl select-none"
              style="transform: translate({panX}px, {panY}px) scale({scale}); transition: {isDragging ? 'none' : 'transform 0.15s ease-out'};"
              draggable="false"
            />
          </div>
        {:else}
          <div class="flex flex-col items-center gap-2 text-white/40">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1">
              <rect x="3" y="3" width="18" height="18" rx="2"/><circle cx="8.5" cy="8.5" r="1.5"/><polyline points="21 15 16 10 5 21"/>
            </svg>
            <span class="text-sm">Could not load image</span>
          </div>
        {/if}
      {/await}
    {/if}
  </div>

  {#if index < assets.length - 1}
    <button
      class="absolute right-3 top-1/2 z-10 -translate-y-1/2 rounded-full bg-black/50 p-3 text-white transition hover:bg-white/20"
      on:click={next} title="Next (→)"
    >
      <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
        <polyline points="9 18 15 12 9 6"/>
      </svg>
    </button>
  {/if}

  <div class="absolute bottom-0 left-0 right-0 flex items-center justify-between bg-gradient-to-t from-black/70 to-transparent px-6 py-4">
    <div class="min-w-0">
      <p class="truncate text-sm font-medium text-white">{filename}</p>
      {#if dateStr}<p class="text-xs text-white/60">{dateStr}</p>{/if}
    </div>
    <div class="ml-4 flex flex-shrink-0 items-center gap-3">
      {#if !isVideo}
        <div class="flex items-center gap-1 rounded bg-black/50 px-1 py-0.5">
          <button
            class="rounded p-1 text-white/60 transition hover:text-white hover:bg-white/10"
            on:click={() => zoomBy(1/1.25)} title="Zoom out (−)"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
              <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/><line x1="8" y1="11" x2="14" y2="11"/>
            </svg>
          </button>
          <button
            class="min-w-[2.5rem] text-center text-xs text-white/50 hover:text-white transition"
            on:click={resetZoom} title="Reset zoom (0)"
          >{zoomPct}%</button>
          <button
            class="rounded p-1 text-white/60 transition hover:text-white hover:bg-white/10"
            on:click={() => zoomBy(1.25)} title="Zoom in (+)"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-3.5 w-3.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5">
              <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/><line x1="11" y1="8" x2="11" y2="14"/><line x1="8" y1="11" x2="14" y2="11"/>
            </svg>
          </button>
        </div>
      {/if}
      {#if isVideo}
        <span class="flex items-center gap-1 rounded bg-white/10 px-2 py-0.5 text-xs text-white/70">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3" viewBox="0 0 24 24" fill="currentColor"><polygon points="5 3 19 12 5 21 5 3"/></svg>
          VIDEO
        </span>
      {/if}
      <span class="text-xs text-white/40">{index + 1} / {assets.length}</span>
    </div>
  </div>
</div>
