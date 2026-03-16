<script lang="ts">
  import { uploads } from '../stores/uploads';

  $: folders = $uploads.folders;

  async function remove(path: string) {
    await uploads.removeFolder(path);
  }
</script>

{#if folders.length === 0}
  <p class="text-sm text-[#585b70]">No folders are being watched. Add one above.</p>
{:else}
  <ul class="space-y-2">
    {#each folders as folder}
      <li class="flex items-center justify-between rounded-lg bg-[#1e1e2e] px-4 py-2">
        <div class="flex items-center gap-2 overflow-hidden">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 flex-shrink-0 text-[#4250af]" viewBox="0 0 24 24" fill="currentColor">
            <path d="M10 4H4a2 2 0 00-2 2v12a2 2 0 002 2h16a2 2 0 002-2V8a2 2 0 00-2-2h-8l-2-2z"/>
          </svg>
          <span class="truncate text-sm text-[#cdd6f4]" title={folder}>{folder}</span>
        </div>
        <button
          on:click={() => remove(folder)}
          class="ml-3 flex-shrink-0 rounded p-1 text-[#585b70] transition hover:bg-[#45475a] hover:text-red-400"
          title="Remove folder"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
          </svg>
        </button>
      </li>
    {/each}
  </ul>
{/if}
