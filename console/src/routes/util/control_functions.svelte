<script context="module">

  export function setButtonClassIsSuccess(setIndex, buttonArray) {
    for (let index = 0; index < buttonArray.length; index++) {
      if (index == setIndex) {
        buttonArray[index] = 'button is-success'
        continue
      }
      buttonArray[index] = 'button'
    }
    return buttonArray
  }

</script>