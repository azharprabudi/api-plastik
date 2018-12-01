<h3>Hi !! this is my first time, using golang for development app.</h3>

<b>Here the rules i want to apply to this repo :</b>

<ol>
<li>
  <b>Give domain prefix at function, to create a new instance</b>
  <pre>
    <code>
func NewSellerService() {} 
    </code>
  </pre>
</li>

<li>
  <b>Give domain prefix at variable name</b>
  <pre>
    <code>
sellerID := "examples"
    </code>
  </pre>
</li>

<li>
  <b>If there is one dependency to service, or another just give service name dont give a namespace. (but for domain, just give a service name)</b>
  <pre>
    <code>
(GOOD)
type SellerPresentation struct {
  service: "xxxx"
}
(BAD)
type SellerPresentation struct {
  sellerService: "xxxx" // if there is one service dont give a namespace, but if more then give it
}
    </code>
  </pre>
</li>

<li>
  <b>Dont give separator at filename</b>
</li>

<li>
  <b>Dont forget to create testing unit (later)</b>
</li>
</ol>
