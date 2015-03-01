# Graphing Objects

![](/ipfs/QmbefthRKDReojALJi8nGPwvUVPqe1aXdoD9ysX44aUfvG/graph.png)

When using ipfs for storing files, or writing more complex datastructures,
it is often very useful to visualize the merkledag being created. For this,
I wrote a simple tool called `graphmd` (graph merkle dag).

`graphmd` is a very short shell script ([source](./graphmd)). It uses the
`ipfs refs --format` flag to produce `dot` output.

## Install graphmd

`graphmd` will be in its own repo soon, but for now you can install it with:

```
ipfs cat Qmcd7Sebd46vxDWjbUERK8w82zp8sgWTtHT5c93kzr2v3M  >/usr/local/bin/graphmd
chmod +x /usr/local/bin/graphmd
```

## graphmd usage

```
> graphmd
usage: graphmd <ipfs-path>...
output merkledag links in graphviz dot

use it with dot:
  bin/graphmd QmZPAMWUfLD95GsdorXt9hH7aVrarb2SuLDMVVe6gABYmx | dot -Tsvg
  bin/graphmd QmZPAMWUfLD95GsdorXt9hH7aVrarb2SuLDMVVe6gABYmx | dot -Tpng
  bin/graphmd QmZPAMWUfLD95GsdorXt9hH7aVrarb2SuLDMVVe6gABYmx | dot -Tpdf
```

## Example

Given this [demo](/ipfs/QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF) directory:

```sh
> tree demo
demo
├── cat.jpg
└── test
    ├── bar
    ├── baz
    │   ├── b
    │   └── f
    └── foo

2 directories, 5 files
```

Add the files to ipfs

```sh
> ipfs add -r demo
added QmajFHHivh25Qb2cNbnnnEeUe1gDLHX9ta7hs2XKX1vazb demo/cat.jpg
added QmTz3oc4gdpRMKP2sdGUPZTAGRngqjsi99BPoztyP53JMM demo/test/bar
added QmTz3oc4gdpRMKP2sdGUPZTAGRngqjsi99BPoztyP53JMM demo/test/baz/b
added QmYNmQKp6SuaVrpgWRsPTgCQCnpxUYGq76YEKBXuj2N4H6 demo/test/baz/f
added QmX1ebVUtfY11ZCpVmqyE5mDoN62SpLd8eLPpg5GGV1ABt demo/test/baz
added QmYNmQKp6SuaVrpgWRsPTgCQCnpxUYGq76YEKBXuj2N4H6 demo/test/foo
added QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv demo/test
added QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF demo
```

Use `graphmd` to generate the dot output

```sh
> graphmd QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF >graph.dot
digraph {
  graph [rankdir=LR];
  QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF [fontsize=8 shape=box];
  QmajFHHivh25Qb2cNbnnnEeUe1gDLHX9ta7hs2XKX1vazb [fontsize=8 shape=box];
  QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF -> QmajFHHivh25Qb2cNbnnnEeUe1gDLHX9ta7hs2XKX1vazb [label="cat.jpg"];
  QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF [fontsize=8 shape=box];
  QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv [fontsize=8 shape=box];
  QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF -> QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv [label="test"];
  QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv [fontsize=8 shape=box];
  QmTz3oc4gdpRMKP2sdGUPZTAGRngqjsi99BPoztyP53JMM [fontsize=8 shape=box];
  QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv -> QmTz3oc4gdpRMKP2sdGUPZTAGRngqjsi99BPoztyP53JMM [label="bar"];
  QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv [fontsize=8 shape=box];
  QmX1ebVUtfY11ZCpVmqyE5mDoN62SpLd8eLPpg5GGV1ABt [fontsize=8 shape=box];
  QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv -> QmX1ebVUtfY11ZCpVmqyE5mDoN62SpLd8eLPpg5GGV1ABt [label="baz"];
  QmX1ebVUtfY11ZCpVmqyE5mDoN62SpLd8eLPpg5GGV1ABt [fontsize=8 shape=box];
  QmTz3oc4gdpRMKP2sdGUPZTAGRngqjsi99BPoztyP53JMM [fontsize=8 shape=box];
  QmX1ebVUtfY11ZCpVmqyE5mDoN62SpLd8eLPpg5GGV1ABt -> QmTz3oc4gdpRMKP2sdGUPZTAGRngqjsi99BPoztyP53JMM [label="b"];
  QmX1ebVUtfY11ZCpVmqyE5mDoN62SpLd8eLPpg5GGV1ABt [fontsize=8 shape=box];
  QmYNmQKp6SuaVrpgWRsPTgCQCnpxUYGq76YEKBXuj2N4H6 [fontsize=8 shape=box];
  QmX1ebVUtfY11ZCpVmqyE5mDoN62SpLd8eLPpg5GGV1ABt -> QmYNmQKp6SuaVrpgWRsPTgCQCnpxUYGq76YEKBXuj2N4H6 [label="f"];
  QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv [fontsize=8 shape=box];
  QmYNmQKp6SuaVrpgWRsPTgCQCnpxUYGq76YEKBXuj2N4H6 [fontsize=8 shape=box];
  QmNtpA5TBNqHrKf3cLQ1AiUKXiE4JmUodbG5gXrajg8wdv -> QmYNmQKp6SuaVrpgWRsPTgCQCnpxUYGq76YEKBXuj2N4H6 [label="foo"];
}
```

Pipe it to `dot` to produce `svg`, `pdf`, `png` or [whatever](http://www.graphviz.org/Documentation/dotguide.pdf)

```
graphmd QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF | dot -Tsvg >output/graph.svg
graphmd QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF | dot -Tpdf >output/graph.pdf
graphmd QmRCJXG7HSmprrYwDrK1GctXHgbV7EYpVcJPQPwevoQuqF | dot -Tpng >output/graph.png
```

Et voilà: [svg](/ipfs/QmbefthRKDReojALJi8nGPwvUVPqe1aXdoD9ysX44aUfvG/graph.svg), [pdf](/ipfs/QmbefthRKDReojALJi8nGPwvUVPqe1aXdoD9ysX44aUfvG/graph.pdf), [png](/ipfs/QmbefthRKDReojALJi8nGPwvUVPqe1aXdoD9ysX44aUfvG/graph.png)

![](/ipfs/QmbefthRKDReojALJi8nGPwvUVPqe1aXdoD9ysX44aUfvG/graph.png)

## File blocks

`graphmd` is particularly useful to inspect file blocking algorithms.
For example, here is what the `ipfs` binary looks like with the default
semi-balanced indirect block chunking:

- [dot output](/ipfs/QmQ8yWC1SGn73P1SPSw8iqSBGEscve1N6sQpzd1xzD5EV1/graph.dot)
- [svg render](/ipfs/QmQ8yWC1SGn73P1SPSw8iqSBGEscve1N6sQpzd1xzD5EV1/graph.svg)
- [pdf render](/ipfs/QmQ8yWC1SGn73P1SPSw8iqSBGEscve1N6sQpzd1xzD5EV1/graph.pdf)

![](/ipfs/QmQ8yWC1SGn73P1SPSw8iqSBGEscve1N6sQpzd1xzD5EV1/graph.svg)

by [Juan Benet](https://github.com/jbenet)
