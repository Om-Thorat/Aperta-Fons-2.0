<script lang="ts">
  import Hero from "./lib/Hero.svelte";
  import Timer from "./lib/Timer.svelte";

</script>

<main class="bg-black h-[100svh] flex flex-col items-center gap-8">
  <Hero />
  <Timer final="31 dec 2023" />
</main>
