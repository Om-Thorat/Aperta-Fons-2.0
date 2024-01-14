<script lang="ts">
    let user:string;
    let dialog:HTMLDialogElement;

    function submit(){
        localStorage.setItem("user",user);
        dialog.close()
    }

</script>



<dialog id="login" class="modal" bind:this={dialog}>
    <div class="modal-box bg-transparent outline outline-2 backdrop-blur-md">
      <h3 class="font-bold text-lg font-mono tracking-wider">Input github username!</h3>
      <input class="my-4 input input-bordered focus-within:outline-none w-3/5" bind:value={user}>
      <button class="btn btn-outline btn-neutral hover:btn-success h-2/5 mx-3 w-1/5" on:click={submit}>login</button>
    </div>
  </dialog>