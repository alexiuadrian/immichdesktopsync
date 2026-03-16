<script lang="ts">
  import { auth } from '../stores/auth';

  let serverUrl = '';
  let email = '';
  let password = '';
  let submitting = false;
  let errorMsg = '';

  async function handleSubmit() {
    errorMsg = '';
    submitting = true;
    try {
      await auth.login(serverUrl.trim(), email.trim(), password);
    } catch (e: any) {
      errorMsg = String(e).replace(/^Error: /, '');
    } finally {
      submitting = false;
    }
  }
</script>

<div class="flex h-full items-center justify-center bg-[#1e1e2e]">
  <div class="w-full max-w-md rounded-2xl bg-[#313244] p-8 shadow-2xl">
    <!-- Logo -->
    <div class="mb-8 text-center">
      <div class="mx-auto mb-3 flex h-16 w-16 items-center justify-center rounded-2xl bg-[#4250af]">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-9 w-9 text-white" viewBox="0 0 24 24" fill="currentColor">
          <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-1 14H9V8h2v8zm4 0h-2V8h2v8z"/>
        </svg>
      </div>
      <h1 class="text-2xl font-bold text-[#cdd6f4]">Immich Desktop Sync</h1>
      <p class="mt-1 text-sm text-[#a6adc8]">Connect to your Immich server</p>
    </div>

    <form on:submit|preventDefault={handleSubmit} class="space-y-4">
      <div>
        <label class="mb-1 block text-sm font-medium text-[#a6adc8]" for="server">Server URL</label>
        <input
          id="server"
          type="url"
          bind:value={serverUrl}
          placeholder="https://photos.example.com"
          required
          class="w-full rounded-lg border border-[#45475a] bg-[#1e1e2e] px-4 py-2.5 text-[#cdd6f4] placeholder-[#585b70] outline-none focus:border-[#4250af] focus:ring-2 focus:ring-[#4250af]/30 transition"
        />
      </div>

      <div>
        <label class="mb-1 block text-sm font-medium text-[#a6adc8]" for="email">Email</label>
        <input
          id="email"
          type="email"
          bind:value={email}
          placeholder="user@example.com"
          required
          class="w-full rounded-lg border border-[#45475a] bg-[#1e1e2e] px-4 py-2.5 text-[#cdd6f4] placeholder-[#585b70] outline-none focus:border-[#4250af] focus:ring-2 focus:ring-[#4250af]/30 transition"
        />
      </div>

      <div>
        <label class="mb-1 block text-sm font-medium text-[#a6adc8]" for="password">Password</label>
        <input
          id="password"
          type="password"
          bind:value={password}
          placeholder="••••••••"
          required
          class="w-full rounded-lg border border-[#45475a] bg-[#1e1e2e] px-4 py-2.5 text-[#cdd6f4] placeholder-[#585b70] outline-none focus:border-[#4250af] focus:ring-2 focus:ring-[#4250af]/30 transition"
        />
      </div>

      {#if errorMsg}
        <div class="rounded-lg bg-red-900/40 border border-red-500/50 px-4 py-2.5 text-sm text-red-300">
          {errorMsg}
        </div>
      {/if}

      <button
        type="submit"
        disabled={submitting}
        class="mt-2 w-full rounded-lg bg-[#4250af] px-4 py-2.5 font-semibold text-white transition hover:bg-[#5c6bc0] disabled:cursor-not-allowed disabled:opacity-60"
      >
        {submitting ? 'Signing in…' : 'Sign in'}
      </button>
    </form>
  </div>
</div>
