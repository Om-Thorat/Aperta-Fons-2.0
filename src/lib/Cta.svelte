<script lang="ts">
    import { state } from '../store';
    function dash(){
        state.set("dashboard")
    }
</script>


<button class="btn btn-outline h-3/5 w-1/3" on:click={dash}>Dashboard</button>
<button class="btn btn-outline h-3/5 w-1/3">Discord</button>
