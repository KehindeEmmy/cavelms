<script lang="ts">
  import { browser } from "$app/environment";
  import { enhance } from "$app/forms";
  import { page } from "$app/stores";

  export let form: FormData;
  let loading = false;
  let email = $page.url.searchParams.get("email");
  $: console.log(form);
  $: if (form && browser) email = form?.email ?? email;
</script>

<div class="hero min-h-screen bg-base-200">
  <div class="hero-content text-center">
    <div class="max-w-md">
      <h1 class="text-5xl font-bold">Verify Your Email</h1>
      <p class="py-6">
        We have sent an email to
        <span class="text-primary">{email}</span>
        pelase follow a link to verify your email.
      </p>
      {#if !email}
        <a href="/login" class="btn btn-primary mb-4">Get Started</a>
      {/if}
      <div class="flex justify-center space-x-2">
        <span> Did’t receive an email?</span>

        <form
          action="?/resend"
          method="post"
          use:enhance={() => {
            loading = true;
            return ({ result, update }) => {
              update();
            };
          }}
        >
          <input hidden name="email" bind:value={email} type="text" />
          <button class="btn-ghost text-primary">
            {#if loading}
              <span class="loading loading-spinner loading-sm" />
            {:else}
              Resend
            {/if}
          </button>
        </form>
      </div>
    </div>
  </div>
</div>
