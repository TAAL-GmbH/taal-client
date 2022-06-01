<script>
  // FontAwesome icon...
  import Fa from 'svelte-fa'
  import { faFile, faUpload } from '@fortawesome/free-solid-svg-icons'

  let dropZone

  let files

  $: if (files) {
    const a = []

    readFile(files[0])
      .then((data) => {
        a.push({
          file: files[0],
          data,
        })
        arr = Object.assign([], a)
      })
      .catch((err) => {
        a.push({
          file: files[0],
          error: err,
        })
        arr = Object.assign([], a)
      })
  }

  let arr = []

  $: console.log('aRR:', arr)

  let over = false

  function handleDragEnter(e) {
    e.preventDefault()
    over = true
  }

  function handleDragLeave(e) {
    e.preventDefault()
    over = false
  }

  async function handleDragDrop(e) {
    e.preventDefault()

    const data = e.dataTransfer
    const fileList = await data.files

    const a = []

    for (let i = 0; i < fileList.length; i++) {
      try {
        const data = await readFile(fileList[i])
        a.push({
          file: fileList[i],
          data,
        })
      } catch (err) {
        a.push({
          file: fileList[i],
          error: err,
        })
      }
    }

    arr = Object.assign([], a)
  }

  function readFile(file) {
    return new Promise((resolve, reject) => {
      let reader = new FileReader()

      reader.onload = () => {
        const data = reader.result
        let view = new Uint8Array(data)
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
    class="column is-8 droparea {over ? 'over' : ''}"
    on:dragenter={handleDragEnter}
    on:dragleave={handleDragLeave}
    on:drop={handleDragDrop}
    bind:this={dropZone}
    id="drop_zone"
    ondragover="return false"
  >
    <Fa icon={faFile} size="3x" />
    <p class="top">
      Upload something with the file dialog or by dragging and dropping files
      onto the dashed region
    </p>
    <p class="control is-small">Only certain types will be allowed.</p>

    <label class="file-label top">
      <input
        class="file-input"
        type="file"
        id="file"
        name="file"
        capture
        accept="image/*, audio/*, application/json, application/pdf, video/*, text/*"
        bind:files
      />
      <span class="file-cta">
        <span class="file-icon">
          <Fa icon={faUpload} />
        </span>
        <span class="file-label"> Choose a file… </span>
      </span>
    </label>
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

  .top {
    padding-top: 1rem;
  }

  .over {
    border-color: red;
  }
</style>
