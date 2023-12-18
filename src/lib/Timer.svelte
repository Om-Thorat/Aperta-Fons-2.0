<script lang="ts">
  import { onMount } from "svelte";
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

<div class="flex flex-col gap-3 items-start w-[50%]">
    <span class="text-lg">The contest is live until</span>
    <span class="text-5xl font-mono font-extrabold">
        {days}:{hours}:{mins}:{secs}
    </span>
</div>