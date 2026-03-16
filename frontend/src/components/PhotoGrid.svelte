<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { getThumbUrl } from '../lib/thumbCache';

  export interface Asset {
    id: string;
    type: string;
    originalPath: string;
    localDateTime?: string;
    fileCreatedAt?: string;
  }

  export let assets: Asset[] = [];

  const dispatch = createEventDispatcher<{ select: number }>();

  const MAX_CONCURRENT = 6;
  let active = 0;
  const waitQueue: Array<() => void> = [];

  function acquire(): Promise<void> {
    return new Promise(resolve => {
      if (active < MAX_CONCURRENT) { active++; resolve(); }
      else waitQueue.push(() => { active++; resolve(); });
    });
  }

  function release() {
    active--;
    const next = waitQueue.shift();
    if (next) next();
  }

  async function loadThumb(assetId: string): Promise<string> {
    await acquire();
    try { return await getThumbUrl(assetId); }
    finally { release(); }
  }

  let thumbPromises: Record<string, Promise<string>> = {};

  function lazyThumb(node: HTMLElement, assetId: string) {
    const io = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting) {
          io.disconnect();
          if (!(assetId in thumbPromises)) {
            thumbPromises = { ...thumbPromises, [assetId]: loadThumb(assetId) };
          }
        }
      },
      { rootMargin: '200px' }
    );
    io.observe(node);
    return { destroy() { io.disconnect(); } };
  }
</script>

<div class="grid grid-cols-3 gap-1 sm:grid-cols-4 md:grid-cols-5 lg:grid-cols-6 xl:grid-cols-8">
  {#each assets as asset, i (asset.id)}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div
      class="group relative aspect-square cursor-pointer overflow-hidden rounded bg-[#313244]"
      use:lazyThumb={asset.id}
      on:click={() => dispatch('select', i)}
    >
      {#if asset.id in thumbPromises}
        {#await thumbPromises[asset.id]}
          <div class="h-full w-full animate-pulse bg-[#45475a]" />
        {:then url}
          {#if url}
            <img
              src={url}
              alt="thumbnail"
              class="h-full w-full object-cover transition duration-200 group-hover:scale-105 group-hover:brightness-90"
            />
          {:else}
            <div class="flex h-full w-full items-center justify-center text-[#585b70]">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                <rect x="3" y="3" width="18" height="18" rx="2"/>
                <circle cx="8.5" cy="8.5" r="1.5"/>
                <polyline points="21 15 16 10 5 21"/>
              </svg>
            </div>
          {/if}
        {:catch}
          <div class="flex h-full w-full items-center justify-center text-[#585b70]">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
              <rect x="3" y="3" width="18" height="18" rx="2"/>
              <circle cx="8.5" cy="8.5" r="1.5"/>
              <polyline points="21 15 16 10 5 21"/>
            </svg>
          </div>
        {/await}
      {:else}
        <div class="h-full w-full bg-[#313244]" />
      {/if}

      {#if asset.type === 'VIDEO'}
        <div class="absolute bottom-1 right-1 rounded bg-black/60 p-0.5">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 text-white" viewBox="0 0 24 24" fill="currentColor">
            <path d="M8 5v14l11-7z"/>
          </svg>
        </div>
      {/if}
    </div>
  {/each}
</div>
