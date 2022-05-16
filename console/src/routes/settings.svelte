<div class="field">
  <label for="text1" class="label">Name</label>
  <div id="text1" class="control">
    <input class="input" type="text" placeholder="Text input" />
  </div>
</div>

<div class="field">
  <label class="label" for="email">Email</label>
  <div class="control has-icons-left has-icons-right">
    <input
      id="email"
      class="input is-danger"
      type="email"
      placeholder="Email input"
      value="hello@"
    />
    <span class="icon is-small is-left">
      <i class="fas fa-envelope" />
    </span>
    <span class="icon is-small is-right">
      <i class="fas fa-exclamation-triangle" />
    </span>
  </div>
  <p class="help is-danger">This email is invalid</p>
</div>

<div class="field">
  <label class="label" for="drop">Subject</label>
  <div class="control">
    <div id="drop" class="select">
      <select>
        <option>Select dropdown</option>
        <option>With options</option>
      </select>
    </div>
  </div>
</div>

<div class="field">
  <label class="label" for="msg">Message</label>
  <div class="control">
    <textarea id="msg" class="textarea" placeholder="Textarea" />
  </div>
</div>

<div class="field">
  <div class="control">
    <label class="checkbox">
      <input type="checkbox" />
      I agree to the <a href="#/">terms and conditions</a>
    </label>
  </div>
</div>

<div class="field">
  <div class="control">
    <label class="radio">
      <input type="radio" name="question" />
      Yes
    </label>
    <label class="radio">
      <input type="radio" name="question" />
      No
    </label>
  </div>
</div>

<div class="field is-grouped">
  <div class="control">
    <button class="button is-link">Submit</button>
  </div>
  <div class="control">
    <button class="button is-link is-light">Cancel</button>
  </div>
</div>
