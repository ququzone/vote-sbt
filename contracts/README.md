## Contracts

### Deploy

```
forge script script/Deploy.s.sol:Deploy --broadcast -vvvv --legacy \
  --rpc-url $ETH_RPC_URL --private-key $PRIVATE_KEY

forge create --legacy --rpc-url $ETH_RPC_URL \
  --constructor-args "https://nft.iotex.io/tokens/ivoted/" \
  --private-key $PRIVATE_KEY src/VoteSBT.sol:VoteSBT
```
