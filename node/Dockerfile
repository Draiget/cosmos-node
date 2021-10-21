FROM golang:latest AS build-4

WORKDIR /gaia
RUN git clone -b v4.2.1 https://github.com/cosmos/gaia.git .
RUN make install
RUN gaiad init chooseanicehandle
RUN wget https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz -O genesis.cosmoshub.json.gz

FROM golang:latest AS build-5

WORKDIR /gaia
RUN git clone -b v5.0.8 https://github.com/cosmos/gaia.git .
RUN make install

FROM scratch AS export
COPY --from=build-4 /gaia/genesis.cosmoshub.json.gz ./content/
COPY --from=build-4 /go/bin/gaiad ./content/gaiad4
COPY --from=build-5 /go/bin/gaiad ./content/gaiad5

#RUN gaiad init chooseanicehandle
#
#wget https://github.com/cosmos/mainnet/raw/master/genesis.cosmoshub-4.json.gz
#gzip -d genesis.cosmoshub-4.json.gz
#mv genesis.cosmoshub-4.json ~/.gaia/config/genesis.json
#gaiad start --p2p.seeds bf8328b66dceb4987e5cd94430af66045e59899f@public-seed.cosmos.vitwit.com:26656,cfd785a4224c7940e9a10f6c1ab24c343e923bec@164.68.107.188:26656,d72b3011ed46d783e369fdf8ae2055b99a1e5074@173.249.50.25:26656,ba3bacc714817218562f743178228f23678b2873@public-seed-node.cosmoshub.certus.one:26656,3c7cad4154967a294b3ba1cc752e40e8779640ad@84.201.128.115:26656 --x-crisis-skip-assert-invariants