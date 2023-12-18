<script lang="ts">
  import { onMount } from "svelte";
  import Cta from "./Cta.svelte";
    let days = "99";
    let hours = "99";
    let mins = "99";
    let secs = "99";
    import { sleep } from "../scripts/utils";
    export let final:string;
    async function countdown(){
        const end = new Date(final)
        while (true){
            const date = new Date()
            var diffdate:Date = new Date(end - date)
            days = diffdate.getDate().toString().padStart(2,'0')
            hours = diffdate.getHours().toString().padStart(2,'0')
            mins = diffdate.getMinutes().toString().padStart(2,'0')
            secs = diffdate.getSeconds().toString().padStart(2,'0')
            await sleep(1);
        }
    }
    onMount(countdown)
</script>

<div class="flex flex-row gap-3 w-[50%] items-center">
    <div class="w-1/2 h-full flex flex-row gap-7 justify-start items-end">
        <Cta />
    </div>
    <div class="w-1/2 flex flex-col gap-3 items-end mr-10">
    <span class="text-lg">The contest is live until</span>
    <span class="text-5xl font-mono font-extrabold">
        {days}:{hours}:{mins}:{secs}
    </span>
    </div>
</div>