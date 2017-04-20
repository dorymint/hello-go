gomem
=====
1. interactive
2. accept in/out JSON

---
**mermaid**
```mermaid
graph TD

subgraph CLI
  out[stdout or file]
  in[stdin]
end

subgraph JSON
  JSON(STRUCTURE<br>name:<br>content:<br>date:<br>tag:)
end

subgraph gomem
  subgraph METHODS
    method[get<br>push<br>rm...etc]
  end

  subgraph INTERACTIVE
    loop(loop: accept commands)
    p
  end

  subgraph JSON cache
    cache
  end

  loop---method
  cache---method
end

p-->puite
method---JSON
in---loop
out---method
```
