<script>
  let dropZone

  let over = false

  function handleDragEnter(e) {
    over = true
    console.log('You are dragging over the ' + e.target.getAttribute('id'))
  }

  function handleDragLeave(e) {
    over = false
    console.log('You left the ' + e.target.getAttribute('id'))
  }

  async function handleDragDrop(e) {
    e.preventDefault()

    const data = e.dataTransfer

    console.log('DATA:', data)

    const files = data.files
    console.log('FILES:', files)

    const arr = [...files].map(async (file) => {
      console.log(file)

      const data = await readFile(file)
      return {
        file,
        data,
      }
    })
    // for (let i = 0; i < files.length; i++) {
    //   console.log(files[i])

    //   readFile(files[i])
    //     .then((data) => {
    //       arr.push({
    //         file: files[i],
    //         data,
    //       })
    //     })
    //     .catch((err) => {
    //       arr.push({
    //         file: files[i],
    //         error: err,
    //       })
    //     })
    // }

    console.log('ARRAY:', arr)
  }

  function readFile(file) {
    return new Promise((resolve, reject) => {
      let reader = new FileReader()

      reader.onload = () => {
        const data = reader.result
        let view = new Uint8Array(data)
        console.log(new TextDecoder('utf-8').decode(view))
        resolve(view)
      }

      reader.onerror = () => {
        reject(reader.error)
      }

      reader.readAsArrayBuffer(file)
    })
  }
</script>

<section class="columns is-mobile is-centered">
  <div
    class="column is-8 droparea {over === true ? 'over' : ''}"
    on:dragenter={handleDragEnter}
    on:dragleave={handleDragLeave}
    on:drop={handleDragDrop}
    bind:this={dropZone}
    id="drop_zone"
    ondragover="return false"
  >
    <i class="fas fa-images fa-3x" />
    <p>
      Upload something with the file dialog or by dragging and dropping files
      onto the dashed region
    </p>
    <p class="control is-small">Only certain types will be allowed.</p>

    <!-- <div class="field is-grouped is-grouped-left">
      <div class="control">
        <button class="button is-primary" on:click={writeData}
          >Submit transaction</button
        >
      </div>
      <div class="control">
        <button class="button is-light" on:click={reset}>Reset</button>
      </div>
    </div> -->
  </div>
</section>

<style>
  .droparea {
    margin: 1rem auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    /* width: 384px; */
    /* max-width: 100%; */
    /* height: 160px; */
    border: 2px dashed grey;
    border-radius: 15px;
  }

  i {
    flex-grow: 1;
    padding-top: 1rem;
  }

  .over {
    border-color: red;
  }
</style>
