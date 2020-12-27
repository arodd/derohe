// Copyright 2017-2018 DERO Project. All rights reserved.
// Use of this source code in any form is governed by RESEARCH license.
// license can be found in the LICENSE file.
// GPG: 0F39 E425 8C65 3947 702A  8234 08B2 0360 A03A 9DE8
//
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
// STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF
// THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package regpool

//import "fmt"
//import "bytes"
import "testing"
import "encoding/hex"

import log "github.com/sirupsen/logrus"

import "github.com/deroproject/derosuite/globals"
import "github.com/deroproject/derosuite/transaction"

// test the mempool interface with valid TX
func Test_mempool_Full_block_72_2c4738d3856e8e3e8f9fc4818a9197d4911af3010e067ec56d08c264627cb547(t *testing.T) {

	// this tx is from block 72
	// tx_id 2c4738d3856e8e3e8f9fc4818a9197d4911af3010e067ec56d08c264627cb547
	// it is a ringct full txt
	tx_hex := "0200010200050401030301b27b6cf2e0dcfc35f2767608cace34e59b1d507d3ef6c5ad7142a1c3aad0d860020002b9ca87703ad2b72bd3bad724a42d3569ac107c6fa8cae194fbcd560cabc1abd50002d76035fd955ef8b10513f95d11fd7a579bcc1db343be0ddd67d656c505061d0a2101cbd8cdb9b2bd3cc79b720a124c16c5eae988ff2248681dd1db25dc146e8923760180a088b78303b918962e3d710e3fa38942f1b7973d9e5038dd634a1b0013a795ad362dff4c07c8aab3a0983d712894adc5a5ab95e20d27c0a4176b3dfff40abf5914db35ad0372a81534185e2e277d6e97edb519aff7bf15ad5787a69645dbf852243492bc0d23a8c131e046440fe14405725b8c23a4e1b0eaec7599b0e4f95c491f469a7c05dcad515f5c209469702fbd6e5ce3f31a19cf8e357105ab65b041a6b5659fd2e7f43a5705a5c430faa8919a9cac7916a3587b95f90b286fb613c83571164ec345e3ac3fefeba9cb534eb200f1cd414e1ba08b42d572f5683093675c7c7393c30578c243f58df05209153d60ac583c1fec19d34de1c9ac50fc2508ef3a0c101603fe6689822ed4ea38ae22b8ab49b5ce1c1f5f695d79e4fdcd0545b8be5ead280ae4956b38254385d155662b146156ea716003da11facc8069b5b80306394e480c03333d6d8c81464547a3f1554c3b14bc4cab044594ca8f154065f02fd2b91e0d5d1fbd231cad765476bdd145732ebef719000a4547fde18372a06997237a53098f948d034477d034cbb037ce9861dd62994a4c9280e3f6865a9bfb999b2f600e5c2b59b6b576835014b89d095630e2e0bcb2533da55d94459db12e3a6539d70f1159c56d03d32eb7bb692f6ec469987266bd6e700a644b70060b1df9ba1ab405ba58abdcd95d5df5dce80e12c2fb15dbc5fc48d661da51663fec085d4597dd071a17c579d217e1af98b14d5e94ed65255f4ab86b6735749e3e4ff458d930720297ab3d5361646c616999e9d9895b2ecba08802f3df046c81a614483660d4430ff336be4bd999056102feeba4af5caf69e0eb9c8c95ddbba4b77aa1715c10da0e2fdbce23691327c0cf531aa0acbc3c80ca84a242c56200230bb20f39f226650b398515af0eae95bcfa0f5833b45692e00182dec2e4546f49c4e3928273f2790e30da038a905bd4ee0eefc38bb761aa7281ef13b6e664bd9b333c739ed9aeed0783e6987bac3c0f3de463de521ab6f88a50dbbb368f3ce13d463c2f39a674c605f834c4ebb9195626d17c14ed8822abd4946908ebeb097f8f90e4482d12895e0a7308eeddc1efb989d14188abe164ce80a742f1c23621e0f881676e0508a6d10cbf4f40c71c41d47b63705738609fcc6a561899a3a295f91178f61c8f23bc09022fe425b5db9e7191220594cb711b68ec383afdd2073779069ba9d23fe0b53d041a4c95bbed2aedfbc7793189c0f9e1f93d8addc088e9f9c964f8804e383c020513c8ef12e99b90f73b5dbdb6ea82a50e93d74e8e50a07bda15a8fbcdf76e9e09f996d276c9611939cac7b5a3a13c3edec0241b9fcb2c3600146c056519f5ac0473da3a9a4cff0c744b30e18883396822daaf70307c0ee818863f3835a6d9fd0fb6c4cdf98895267a407ae3042d5e422389cac50723a2565724d4de7047d1d703b469bad46b46f9e6ce8e27ef33b00b24570c3560e1443a15d62124b391509f02e85e9c8ebae05ee533c0922b3bd6cc00ccea4167ab7370f615501976553c350b6f7b21f592782476bdf87acab34d53365aa76597ef8db3e4cb2022fa06b75f010ad2d259d139e592bbd52cbe4a9aeb3697c3a5ef726bb86d550c043d2046ef01429ff761fe175632253b7752f2106954c1cb569e513b36452f4d91a93c750001d8e75312c2be5fed07d2dbce8bdf4daa9d5fc161135e83ca61d828565aa3310c4b09004486285c738f7ebc55361f97a05e608e99a2e771c51f559230d264580ee61d8bc4e96ed132ce187be8f069aacc7d72ee216cf311184a1e08daf9a18c0bb07db1ecf5dacdd266ccd4767e2b8ef6c6ad066db45af7d68af29a54ba50cf076ecf21c2e24eef9d27ea635ce0260769c7c76c386a98f2fccda1f52e8495870f2b28149b9f15ef41111c4db005fda31e46625914bab04c5bf293ec0154530f04002bb7fdebc041bcce84ceceb22ee2b3ec7959a32cb6dea38c274b5f6c140009c18f67dbc949f9a4c10b1b3c709335d3991f55bf78613dbb85a9366d52ff7b03468ecc848d2bf68ea3840f40c6621c72c3eda4fdb44261b08bdd8a1ad281e309ca9459d6cb2348ba736b20e55f1c6608c5de6a29523a9575398c9d2bad213304eadc57428c5435cbe8926df2481a31bf0301c69dea0064c18e1ff958e398ec0317d03723cf698698c06403e1f93dea759d801a30f95728db04d3e1a1383feb027da8683e4e5708a5290ff14235d9fda14e9f2c0cae514d058224fd02cdced3067ba2df876ea5f6ba6d2523a53fe33290500c954845fb517173d5ed3b1c86d00fa6066e7dbb523e3c5e61561a003ea95609537eca80b3612ae21c0fe21787f6067a82f20e965062e21403d43d9f4f3dc20a53ed86a72887eb103187ae1ad24504874e565fee47b6679252f35e63fa7ae2521c832f5410fa07ca47bf4a3233ce013ef73d29366c7882a7514d8bb1b6382c587bbd9d4a769fa00f4921afcad9da04d67f8d6795591614b3321a3a45b58c0938102e96ed1ca4c1bbadfb49db33e700b67e84b16e69b6ac24dd5491ef58532abbd513579019df1f9edab7c656686d0df0157db9526307cd81fbfe79a1d3372f4d05946449c88a2a68dfb80522eccb0d661fe7fdd0773fdc67c65839ddb77ff9b3890a52f94e0536620ac3eedfa7940b88d224e7caecbdbc12077d19516f7f3d47625c093d84778a9aff02a8ed4ebe0ad9551ac8f4d59c93bcbd4f5e02686b47f14b48ea8cf15ca4368a946049b6a40b75f28dc053ee8f36a03a51a119c45994b4d7f03422ee437567b1d8269de62104250e6aa8659c6edbf794f60e19c983349ed92edda1fc164f6bd325ba97ab3b02bf83968b74b7c6a9c171ee7798713ccb578ddd7aea14f9215490d98d75ccda03dbd2a38666d7cc8792c80afa931fee64785e1a3156617724f8d7e934fa56c20cc8e9e8965c847794a31b1e5774ac547766ae8603d08022e07bc66f5336139e01eb1b92685611bfd1b907d5d8a72bee15c8f12c2dd685282662860df331a89303c1e2125eeaa08e5bd09ce42b604f6375df5a987080457d0b2d111a82482b030efd2967f096f3e307d114361cf5d88af7d151d573c9800172b7481a95465aff0df8858fa3d4f8c9fc089f90a5766662337e6ab89c37e596917fd92bee8edf430ac71519138f633492c6325aaba4b240cc347479fa26c77138b07fa82b00bc12063710a21290ed2d6d60d9accba8ffc1e4f501c98fb5b4355bc91b0e8788c92f09ce59a51153e0772704212c3d0dc6ec01b9296926e1254acbbff73f25d096e20cd4c86621e609e04df77078e5939aa979be5f2e571b3b422113ddf3b7f11ea704ad61bc954526dfc9bd0ecc339d83446088bebefe1e6dec08709c7dcbd07846069d448776291e202ee003fb51b0c765b10ba7a3c022fce045bd0c04163636b602e06eab18acd86a108bdfd42a8c5ae3314f7726ef35d12d54054402eb32c7740aac15f3a1fd09aaecc13aae4095b03b6fc35a17a62cde59916f07aa5b2abb190d8631490179f472fae6b8ea4b229fe8ed6a8143b5054ccc7b36b7fa0ec6c0110fe064d130b183e70a3ca36927744b8e14dea1d89778601579fb76076b157e2b0515ab0da39da3b2e33d9a153d6006ee8da09b7628066d0f7e1f9207fcf6b68d03085987c906a67c4949e6104444bce8db96481523bb5182e65cefae237c08a80c68fdbba091bbabb038342b330cbfaa62dd3aaac2f39134029c78b368235ff209eb7056b022dc5e6f0bb0e344c310554fd01043a3a0092b446b3dea67f7bde90f6bdfc86a9dd20a83bd0c13f297476b0957fa7b2020bbf258cf0f15ec4af35b0b7fc884c53e1d81e742c73194ee68c47e7d2a17dbcde1599f10cceb75b48f9b0f72fc84d9f359784d69bb0d4f8ff745e6d14421c5b867b69b7468c899cf8e3208673942dcf195da4c7674b4603023cd4e9740fc949a535329144a509084754a09a7d7f93f3930110ea12b1e59936832671ed9c338ed7b12537280db0e3552690625759f2def660f20c30d3b62bca5d10b5a85c2d3a323c428b264f3fca5cda2007f9cc73babbb4113417f195b964e4215e7d06c9f5f2b2984452f5cd0562a660d53b47804be205be45132673a0e87e9b358d188a4bb13f308dfd48e7d3a19cc00cdaee3472b64270d61509025f63d011fea544e4bd406713b3a96189d51926e0a0d4978cac42ac91e75f00af0c556f85c578e9c8cc130a720c8bb691047d4ed0b86a4600f19eeeab92b6c06e3a75f452ce3239f6008db5a20fe95231c9063920d5153d78ba0d47fdfa7b79f6f726bb2161d36bc85136a5f4f4d82b1e6c083b10fded692fe72de12c4984236d1b25bda589712b8d3c66edaed40e94469fdbe700013f834241a6cfb42eaf841d6c4db5b8349b1fd2e56e36b30ea16c56bc5f9bf0e72ffbc2726568acb459f285ce90eca77e2f74bd88cbc6c8d74ed76ea35031d02fd1f129c63e7b2c0b2cfaa223c07321759658ae50360e3e775f0fe71fbeaf405b2a0f552b73d6790869f3436c2f3c736c6a67e6c6dd5a2761e83b88b3399bd02a8103b4b3656b947f2b713dfc453290f0c9bb4d2064e7f51ccf14169d80c0a064f22df85734e06e3055b00d2034e5a5e2cbacb0dc2cc663dad72a638581ea1090ed7a972cfd74e693887a43ba97bdd895109387f530ccf78e1aa4a19a110150f79710c93fe6e630732a6ebc88e0592f042a1bd56f4d8f6ce39e6c1fb2b348e087434bfa3778d205f3058f621711a41681962cdc9d0afe3973714ad80df48960968859d1f47f2366f3d7fb2b62f89715f603475d84e70ee690485dbe579b6300929711cbc966646400f2e6cfc62e370b614427648349634d0bd2ee6149637d3072d2b67859ace2ac967008fd5265aa2b1487a2057ece59e31482ea7b8df517404c2c92f59f55013b79b111b6d193855e5d4eb10e23b816b95644a1e61c7bcc00fda604315d9b884472f051ac13aa7546dfc784b72d7be2dadbdfe7900acc884057becad608b1d379e920f129c5f3a40af6a9d032fc45746f95236dd294efa0c0206fef604277f428f91e15f729a4b073ce12253f9d71c8b528c650c68dab076064b7ec6b0bda314dcb5602b6e892b5f31f032626ea6daebcf00e3f065153afd0c9c107cb4575d14cd35a72dbe72926e1b1986f528344b48b6c6bf817ae0bc9e0fdd9609461fc8cd76c972f028fb462db6c45bceb305812c54788f78602e5e1c01d4b11d6a580e93fab23495e9196eab78e4ff6ebf503e2d9414e8b8e1e5467d02fc1485ff7770e55656b9dcad13e91977379d670246e930d2d1661ec6ef31fd07479c8a9e42533c3328fc9e06bb09daf0e4383ed01394f4ce141c348998f8ec0e232848dd1612e5a8e29c1e50388a52fed2dc454d69ee3feae06d5ed44260530d8b918cc3a38ab589c58f7d855b3b93e5308491b56c6db8459ef7689f6b65d4084948a2c640b270a532d9b9310f600f0e5344d036a62d8f8d070f2c2a581b9a0414ddebe3a2d7afca247a7620ceb4970ec165ea7f0f4b65921ad5b72fdd0e7309b0c5bbc5409047e4a69214c213f6de201fd9705910c271e00d68731e28452604d40e5745098f04af5101d4fcb82351d763b2932807a07910afe31e09f6a199041305c50e8a5ff91464a19da5c06595e4a9398f692cb63dd920cde0a561f618067fa96a1c8aa1bb12012648fda97b16d1f6afe278af426458422c8198ac701f0212b1d43e6fb11c8ada3b3f79c16a6121c9d6fc51125223fc1c2929dea48dc60331d265346fadd531c2a1f7e78fae8c090e8b2353b0f79b58b1974428f7ba2908c6ed42728cfd69ef902800104eae20745161c3b1d07e2313ce4f3f281093c60ff0a47bf7f2006666bbdb4ce7a6e620a546ea61cebef87828d5dfa167b5e509013b9bef499c3e926ca9c0c961920639caa0e630a3b24c8a524f3e64289778030d425399eaba73c14af22f91017982fdfbe0dc118ade0257dfcd4e40643fd9a702c4b72a3044df7b9d0a8cd098e443b62cb6dc7926064a1375c464b372102f050c873d95d1de2d084d97a8ac1fe73f4530caba213c790e3170a7c50b2a460c0602c4a32a44ca512ca7f3edb4f808118cd0af067741985d59dc693ed15383add649552c8c4737e649bc357445a0ac85e38433cac3e9cc45046108d008c04cc3b83b55d04c6434c011f6133d71e109fd088081368f694831e0687530ab7112d0287f20ddf2df27f5646e0cdc2f7a7bfb5d65202dc53a3d735856c895a1d340b64f22638ae26368f32040d627402784993690e829c763ab4d4896f1159f5b2c7d4ffc5461952c57eeeac0d54170000f1eee31452a5c356a32e90af93211322c1c6a617ec29da1d685edbe5322bea3d041c2d8237e47ae977c48199a06e601dee206ec9fb45ea5ea69b6c29deb597e8e848f35d4096e84d0e3cedbb6fce1ad6894b5866acb5eb6bef69833b9bd4328f36becdd1ae503305889f40bf1af3d06eac1517b6ff7ec977b1831efed15da93500e1852eeefc1c99ea82e74f26c5dfb109304738bfe8b667e9ead0df23d1c3536d0afe1b22a907dfd4099aa071ca78a35cdcf0bb61973c8050aa6c1d668207a4c7439ac71110673bc4d2a4f79e87e45b287cf7a156430e07c4669947d47ba861ae2bc77942c25304947021a3efbc120062d9f4466dc59acadab8542c2aa3784d6b92511caa3a16318220beb872d44193e2fd2b8afa174700aba3f9842e621cbd955d4e33c05b2b9a0693fd2dee7c6ff832734926fdac894e6e7fad4602063f502748863f50110a4577d850fc5146b5292bfb1ec4444d51ae9d8f83b7029cf31b763874b3784c4f7a5cc3d40b05223cae404a854cb5cbb4061678c951b20b786709ab5b36add00f2e0bbe6c37c0fa197c805ff50961d7c9a3c30d994b4a489d82dfcf167c7613c57b3197aaf3a3629b4e7fd4759f7d74ce715c908b1dfaf3c13afd7f767b2aa4650bf5cb463ef8b732686a84f14e78e320d77639d72630bf7a7d5c2bba545033b7dfcca089e486977068a6c80017ecf71cb223d8b6eae99b2f857cb600a02fc1487366813199307bcac1ddcace3da6d788ffd94db9a9d674828ae0b900425020531f0bf7ddb470e75568cee0e54f5b875cbadac758f248d07bc98a3a6361c0da6fe4dae978467e59cabf672d5aa56c7005d4b1bf6049069197d4214bea2911966add2467b746569c17e909e2b93f574f1c03f1b2d2ac595cc88e2f839b852e07b53d712b9a028b7405d2dff33adf172be9ad6853c3244b985a49793657d2b4ecf97f5b7aa7c2621379f02f8e7fc2e72dd6df35881f3b8864b06ac00077ff8caf90998858025059398fb9985c9c63632249d006c4aeeae760b7c040174a35334cdef63e8913534091f0f1bdb9cfcbfd7d634e8cc87eb2932ba23de5dd611b538e33b1883bcbfa00b4c228736a632d5eea04fd519daa4be4adb45e17675def3d5dffca0bba0ca8d1293e42087b1c75c90c36f288244a6c99c5d42cb5c9bd02b6999e01321a3ddc2f30e6c2827a63232a5e684e7e2d15c2ed821317f586c89774031a975dd6dc2be0ffdab2476b8e39a087b703cf7013c59a8a9e49a08527622d55336d0442a2fc76fd199be6886aed11e34f55df9abe5e72177bb74bca8d8200c5ad800b7e485c2cf04ff513974f7a67d9df7b79dd2d6a0bb72dcd7d449e7c8a4415a5b18f6a071cccbbe084b97bc5413b9e3b1e629d083859ef560d18dd4a1731f4ba26808a6fc7b6920ce128068181e22566260de418a93d7fc144e5366bd640fcb08b47286a20525fc17a50a1d4f8f01cf31e1ac5930f233e11f16f7cfc0d75c22a26e1b928e1f690677a820990f0d8f0cb725fe6f142179caa96c59c1f7bc7ff92c975a5cba9fb24d388a41aedebe54408b221e9fde3beffc35a65117ce070fb59dad7746bbf6410c990dea68c20697569a770cb11869b15bb0c46bd4c4fde912d54e84f483ac4e13e238c05005ab69e308dcb4e6f5006041b86ea63bc21281f465b3576d5cf62155ade0d535e099dadaf10570b5f2b0389223310bba2a6bc021bbf2e4c741df1644a2405f39e89dc74877738bc74c1216e23f26a4db66d490f67e5cdd9761e0ddb5b47b199ff8c7e4373e1a21502852bbed7de6f06eb3a4bb28f517638d1409d461b7a9ba5dc3e5939e8d828b11f0ee5cb23fae7d06d19e88a7f8584d557a4b0326197ae51ee4a96bcecd79a1151627e53907a77766bda2727ad3fc280049ef5537fbbe96d951a7d34e1f61e9b00942a79ec4645d2f47edb95b1e4182e57a3fd0631118dc78fe901e37604c0042b6bcf23c503cb6867e377230f2790fb38e42bae1e397b655ff56dbaf7b1571ebc1d77a9cc960a87a5ad0bbfc77e7629290d3d4e7c6183296d342a18337353495979ad55cd7d349896fd7e411f147521034dcfde8d0ffec3bd16c31e887909a6a293d74fb3b9e6ade460416dd18e3d55e16d7882323bcdc5b49c835cc7cb50f993bac540c89d96aca7d5ea94dd107de5125e7f9162a9aab9c84c39eb8430952eb971510dc58e81ad25ccacac867a69db300e8098a54fa7b4dbf41a7f251d8263999d29c6a16e4f3fc728db5480db9e10f9ac7109156de29517737f8bdf4af7cc572ee8b9f6994815860af6cafc6c3b67d340513b85d832eab12c13cc50322f687553e945ed00d26f97411b5dc718659f8b8e6b57b93a4a94e51c9b84ba9548047e28999864f9fe8d5c618aefdb06c0e6ccb3db46e17f95d58f45c8db3e26ebec1d73cc6cc4e3e5f177321807896a2cbb98ad30dc16780b1c5de57f7983c869aabe4b46ee3ee6ba03fb72851da3a0e21e1cd1b15f9a09c84f886f5d7468d264b73a716985d82cbda02177562b0256e06f8c1f00310311429ac36d0ec0cfd9e36e929ac1523d79172a192bdca6cb18470f9258d491fde622521bf8975e49123447bc61659bb1433be06de61f02df57a063a40c3e457b62884c02df073e59fd02e2b44d9e33eaeb994c7b626389334577fde5ab351eee0148900c489ffcf80d72c7173ae76747f77b53f9d145b57eea48da0a69572f01fc38c02cbfd13df7a1697e0355b01fe431d090382569401c4a1830451ac205afe2d160a7bdbe44bcbc5161525453e95d4eae803706677a52f0ca98f58f400bb6c23de0aa61bfc621442d39c44178c848a1329e062a95518551ede4b2a29f2275f8eef0e453b1496fe920a2cf8a473fc11e1b83151ab5da3d98ed4f78196276ee4c27f06239c31e080d5efd0a5feebd3543a650075e764fc317ad207e46ac1580841c90744ab4f05958708589f5ead9e13279c87786603ced9091ffef53d038b25480b0061843f35a5305d1bd6b878c475206ba922993859df33da1fe4fef10cfb96cd0083f6b370825d46c691e179373ac82e77c8d888eee9ae96a4a4898a377c57750da0a5b4d3b3763e1a9c588e6de5965aa7f61df595f301e5f993f9a8fa6ecaa70dff59f6f78822a41b276d7f55c6a08558cd5c6d146b2dc64457c7bc1a2a1e360525a55c31c828c031f0f8eb25b6212e2c589abd1ef88b6b39f56387c80a138708040e0bb75781ddf8dced1e3ba2fb1f8dfe49c68d7f81f9a84fdc54c8d95cc406f93f0e772494900d79c52be3cd04fb5022bd5fe467da594ce6450c77cf9b850237b996f03c4568e28422526d30a612d7bdc02864c8134b38d36993295c129403663a763a8e2e540d6eab3dafe641f0a184d3b23c2e6276ebaefc4e2773fedd029a3815b8d6a1129d4f46c74d772d9c279caa4289f53469e34a49a473f6542407f9e0ed855140ee74cd5758d3ae47cc7d713026650c06a11634e0c71415d535063162faa5532b8e23debe40853ee95d08b9f3a855594253a4e1deebcb3b0606005978d410e825b60b19e57c20bbcfb6d0983a49e372afa9d9e713cbd04912850667d75dc98ffee297a102dd5e5dfbc5797fbb8ca078f6a350b16b3b8f6044970b1462b5321e89f97760cd7d98e2e3531df99f460969f9b66e3f10486ff19f6e08d6793e5c0d182a1ac60460ec50b007b79e2c13290f51baef5048c7e081af620f12819aa103bbd1daa7636bfeb494aa6de5d2db1e50115306c9d6aafd8cd5c500b59408c23b5c75b28a5c2148af9f23cb94e0be1d02588bb05346154544a1ed0a5d7ede9fad303a03a8d977cfa25d194d1843d1ad5a33b7efd9d52def5360f300bbea37904625de0b26d918e16e2ad1c8837776190990fba72717fd357b627a04a2fbceff9453d81ee533d45820d1b9ff3b8ed7b0fd23659c69fe901eb7c4d20b849a2c31a937fe6978cf0389a723669814dfd320f33f510b977c07accb0f6a049e29cff30a9cf9684fbc37050463c062b914ac2e390241ec827b05f7d9e28d02924d679292b5b53ffa43cb8f4fc33b05eaf981917c959ffb3d76b4028fa11a0479faf4cb330a46722844a6e36dea14fa4e5229a9ae01f04380febd11ae2af603e449eb841f57746151e3185cb5b2494d959867a4bf80c721c6bd20debac3a1041a29e706e165881e8384a1583b600441d8bf2d36a3d0f1e9bda28187534c6c0f4b90631fdf44a7b647ae39da9285ab494e4fe92132817345bebab4c0aa8ad60e1010fea588648eb15b0fe20c039127ea94f36315a91bfc637f827f2f3b46790c06ab1348bc56d7eaa5563c6f104727fa9a134d988df57d750944e38c6ae3bc0a4be12d4e8c20685cc30808bc5a66a14094d733fe08be61c1335b8f8fbb4190098ed697ed1e0896434add7281ddb1d95a0a266741e630058a61b549c44d6efd0fd9833113e2d0ac4fbe427e41239a61417009bd0992b46e9c30d1431f91d43a0f9fdeaf969185decab082287fd4d69ee1d0fbf3aec85b10bddf09944f6dd5bc09eb2f4e9a1e4c13a71296d21d33729b95beb6d145ed886da6e47d63baa6f20f0287f2c6860911d519025bf9cc5156d6a0909c156479967b2d9dd337337c0986072491d5552264cbdb310c75be84dd7684017d3b2edc4c98b3ed2414b5b5f5580d8be9a913347391c26abebfb67da10ef4fcbe89f1a8535e0dab21b869f72cd40530cc936e0c83d23e240a651ce2f09724f06e190b402ea51609c3a0b83473100b7fd2bbdfd9c6b069185d73d124284cb776c02a748bc6ae70204c3fe8822d5d02086ad8f175f86afca8ebeef8d86fad9abad7019e76b9445008ea8db19e5e1d070abb6a9f3d8db85db66a7f6ced71fa87eed80522a1dd524c1c10c324ad53fc0bbc9846e43408f73408bc14c1322af232702012a6b2a321541785449639362e0891b557e580ed7c9d1a42985bb7e0739bea47b1b554ea562930d9828e68d4c90b2414f7eb8a7dd49c9be8e4f692ce071940bea5d8edd7ed0f2bef20322d0a2b07f621761f5e04b57dc6a3c6bede60d2409be0b1cad4b562e9adb74a2aaada8b0f57f5959cd2711ce263b9900dace4adbbae4f0fa8d1a1d8a5caadeae31da54e042369ebff41cc410a0d2e9f97138b04a310dcf5b2687125b96c579be73dae1f060d07c3283b6c5c6330d395fe50c762ea8fa50150a981543c78afab5ef52d260a6e8a92840e40884f5c96d850dc8f9ba0942b283a330632751bb836fd41c95e0d0c9f53e8a3c08b4a78d028064bad31e6527b06d5191f4c90c380c18b28bb720e81fffaf20e673560bc4b52512e53466a6345e69a1cfe20c420489252e9347a0dc0d74a5dbc4cddd40744dc665f68d046ec32460d727b4303d2d238c91324e7063d08c5cf208787191fae274bb5282cd2a35c9ebc901f86cfd6c75f8e64c6fe01ea3fb9304afdae3d400e9e5f743abba658e997ed7690abcf9ea91af217e8c70f34b15dcd3465d01a66d4ff083d2cd31bc48941ec822b42abcb2029fb91a5c40affb45cc508799039b7cce1cf1da9deb4d05d0596951870d1e25d9df7ae0a47014febcb7d8372ac26edd2b5c22b30bcae9ac062cb13510a36be13173d584ca905131a357a61b1014f4345b151b8e5e92f35b373d550099ce609f85509f34d8c00283c310a6672c3c991cb7be339f3758bc068af8421af960d4e6fe046f9479d063631fb228a2118659cfa11a40e6f1ff14c2dce6398348e99d9abc95e9b33aa058e5f4520a74e7f90a14383a38d224c6e2e25f26ba6a20331ff779cb3351c770fcb9217428693637dc94f5b7174a0a6ad9f15779f089c7215c276d909c644b00814443f24840c9cafde250ee9ed9babae6d54348f1077835666162c16b94a8b045ae832dd4533114673d07c6c1050529948d1ae2c11b11b313bc532475027f10bdadb56cf435e7d7971d2e8f3e753aba45a73b6165351c527aa82dafb5bc96e0de711fbe97c882304d0d0b60d1e016154edbf025b27a9758d3c0714b98547d60c9fdcb582ee35b47ecc8dc7181331c9b1d86dde991731a3ad26a42b053422e105f5b403e6ddadf58066eac7f69e6f0bee96de29cd67960179bf46167685aa95062c044ac38550605635e61a30040803bc5de228fd3c75df292adae7f96e97b50d942f974f84e4261aeb67dd37ca2800a4cae6ac6f89aa687c36462388d06aa808a12d75a3ba0185dc02c41853c470f049f10c166c135603f8dc0cf11b90552c07be245135fb0f2dee11a10d68bcaafd2e5a7b9b0ed419265b9e69c8f8cf67a2074dee2f5387e3678c88511f792cf5003097b849c4cb26c247924dc9ddda75b00cf474ef873af79bb9eb95a01a484aec6e5297f39389f2942a66e1d406c02fd6063d546a1508d383fd6972bbd415e8a2bfce04fd1588a6a4be1a5a32368ddbd80ca86d256374ab90ab120a39c97932242d55745c7e50a8480183a3421228d7050acbfc3075cc333284b2cb108618382e27ded11635054e4325f34f22ad965d040f6efefcb52894245a25111e0d9242acba77866606d21a142853ae684029bd6c0a6f4c1736dffbe610b469f94fa279af7d998fd8d2dd2a15e16b8b9030bdba4502b40505d49b280257cc2617b1cf4119da6363a4737f368eaea3afcd4f26524e086ff827d85dbbc296cd29aec76c4fd16a1ef7dabda71165c44ea7c01ce79e400b00c94e6fc70b250c80a0021e084b63b1e819b1e424e410d290437cda1075eb0ce37b2b7ef24e21850d9b54a19febdd3a3bef5fa338039c7881d0cbeab251e9013b1126d5e4e28d07452ee914de825d797636b4eef0e6d189e10b57eadea8d9007c155340a560c16aa19d167dd6fd040f4331c26f212fa855186bdb3ba7118f0f31f46a3500e59ad379ee9b0f672b0d14c70f796225efcf4c3aa7976b4d55c30d2aa07dce2adbdc8e6086ad23d1c7128f7234e1663bcf8beb3b11dd92802aef050b84617870a2e9273b255c132ca15daabed0a8e5519e533c04e19c43f6ebf30bfee5457e8715563c8cda0908c5f34719952f73f0f47f7365eac3c141c952f40203b51cf3fe5b4a31fe3e277a672d62ac9fa7a959025ab53484202b99a3376d0c8cb1d731f175a32935b2d8132fe5732b6826d808a47da7fcaf26786afb890d0d1016807405fe83a6c66cc32b30e6b9d4dc009f99e2ae2b5f63e8e275c2cd020d5b41d607563f55ffe73814759edc4d9101aa329ecaa072bf7f7bbc49877a5b0e83ef4e00bdb30d25d05879456c23ec567c1da0ef446bd44adfc3e6b0d63fb80c970a123ffada559fdd01ab5f414d08738992f579cb8e2bf26e58f5b3b542700d334c18b84a13de3ebeafeb3d323cccf26b25881cad6f295556eb7c987ee7e20e21be221543f126ef2534bc9377adf19e0e089902fbe1ffb85b61adbf3a1e8b0bf4466b858b5890d8a7cb24e8d08371dfdae2bf938bb38ee6970c8197e8b617049f051268d811cf603d60c627bf62b3ac2b75607d7a5d32ded4d0efd356d3290c183be51fe8013eafda7d1cd7190ffeee31452d5cacbd4450183fac431bfd1d01d91851a8c004a835a2a94d84507f1a780d4c1eb8113273354659557ca09d95054a43d530059137ec28e634a5047f3ce418278541e403e3d7608cc31d9894f50de25b72a4ff702bd76cdb2ed882bcea950d46a0a8682840505b200cbe8311bd0ae16f88984ed477eedf14255cd2d97f33480394aed917f4fe28ce071d5297dc091a0d140e3f99fe443234ed576ae72d350f4d38e67366128a71abe82154fc790b5763d884e9c6b9af192b478729d9f9916962f9c19587026a7395e08a5cb63e0155f894972c6ffa01472bff97ec0caca6fabf87bc28e9e1ffd9462869993a9d078611549c92308373020c9ecb9a3b86944e7b4f199c2dcee420156194be699a06987497eea31cbaadaaec42236ba0d80b6804c78ccbfd971ee496d584d7d3c507f9af01d265d76789df6bbb417a283612e6f178bdfbb5c4ff07bd46f98d2d30001b3dce937b3b13bba358c0c49ac67bc566e6b1126a223d0af1252af3c902450affb12da52b40dc4eb6c593ec6e8dc92210913aa919e2e5662d4bde6d1fd42b0205ee7d7fdc67acfdf2cd67692e4f62cda107e810652f4c97745f556bae6ac3093b1e3d9752ef7e984eaa6372343d29b70dd25bd4f3ea5c08b4d0f31d346e5a0c0445dfc0d41e7955db1eeceec7ff44f06fb84dcb1cf361e512874cad985a9901da3b2f67204f0165e589ebd4e3739df173a7cc8ed08107c4d442121b15b85c0bc0727f6e8073877654aacb4e393a868d6054f413cd899cf4176899facef4ac052cdea604386937c07e4a086ed6be26a7423f9a32c1d7e915426dd8b2083f5004b26d99ad716e7eb55c6d191513e7ecab08c82cd7b29dba583f4f59aa4cabe20242db63f15cfbddb908e643c5e1a99a5d155597968018c643672d05250dd0ab0db36c1cb3060b06c4fdc4ec21623f4fcd1105177dd44afd797dea3d155aa07f91e6cacadc6a2a85d57d6ddc143d46273bf30e6b5158c46d2c4c01ec0e2c4d51b9f09294a9e68bf753060797cce8f2ae2ffb0b793bf3e4bc4c8c88a6061e2a05fac39ce719ff929cee369b09241a16d41ffa091e31e5652467c8e829ae8475e48710fa40065d128314c5bbbcb38431eb468e9f2bf4f706525bb745c47d54c661a1aa371fcce5adbef14daef7a38e64d1cf64d9954df0f9bd9437f347e63888a7586f87b64e8ec576239731707c8f8d8cb93ee931b967b8fd7c8c7e1a320914f63f60bf64f0200b5249d53c47ab22428b1a839b3bb8f0e025781629bf5a388e0b78214d44b81df02cd678fbe17eb5ca68970487d89fccfa4e94e2f7a83cee2d90e9c71f8d270f7b6b562504728f191c782848893be3f1397a68e2206ed3deb2e489555a5db07077175a6f89689f18f62333565824b29fdaef57a8a0d2c44520fd9bfd13d13aa03b342b1224f6ee540c45dc7ea243915d952f4e1334a98543474019650199807f28e9eaa2a2dab2d93a4c84d46e7261a499e923876acd07c834d768e8f6b5b7eb68990b034ebc9cb51e6aaa5edc747b9529b47ce258433a49c6c4754f139521780ef0657f1c32171c6334dfaf1f009c472383878620e2d5c2a61274c2c913dd7d5cfe470c1578088de9c13376da463984e6b3b3d56f1db3b1f5b059dad3d4b663dfeecdf84566d7745b4465a753085e335312ef4f119ba5f54560b8aadfa7b95b370ac5a4ddd55cbb5fdf233043f6af938692e1a634a5a29f190499f4a4914f144ec98fbdf2580564f759771420950ec81f78b486b11fd853f3b014743f2d5b518ceed7b842302b59b6ad313c88c113efeafc91c7291fd56b22507816ef31942e4b575b6f7552a0f3d5fdecf62fd17ced74da69c4c0a81f058b0f70fcd0c5beed103cea702f1f773b72cf20f146281d0d6213632c76f9fdad1123aa61430523a63b212e907056dce796e05886723ef12f16eb546402ce75cbca06c466ed98661551a5cd7b8dd0ed690de21d5dd14fa07914d81fdc662ef00fde7b48a62585895b563a0d350ebc6f574505a136e17888f07955c1889f3a145ebeef02d5d41386357ecdae636c218271b03f5a9f3836c600239870310c77138cda8380697d130866a8bb5f703c83c7fedb0164c0379942990f00c44e6c450388a7bbc3367e7e7adf64cf47d939a02a2498bc6c4a7026746a93f17585b7b72e9832011508e9b1e05c7e98793b96c8274e0df40b7513093841d6d6c019990bd5e38d0512ade72de29e5ca600aef169dc559acbcc3dd61262189ddc33f3684d95e808aa442b5361a798dc93855be1d4c6e79f34ccefca0e4175b2cd9010149cd85331533f5909f188958b3171dd7871f9a9433a22046d73d85f1670d0824ac4bf5867d40d40049d934922d251934a1956e5dc9a91d621f74edacc340d2a1f410998f800756e3f73fbe7f9c63cd2d997bcb5dcdf0b83189be20e0fb0ccfe592c07410b13d725067aa5a86e85b6f8e9abcc5cc95cb309ea83ee90073fa641ed8802a0c69286cd3ce6f1d7962f7d22e1d1888acf0b89e957598cf10054b9fa56da95f4485a308b4ce449ce0d3f95713ad846aa1c0aad321a6d0758e1ba379bd7ea2d6c9b88e3c8f77d53c67d62f16367ab1414175cd4553145db1c18430a0b57892954893fa7eb08d59c010de9895e551c6a0124013a6e275502560c387393d003e119fc31a5390977e87c248caa76fdaf4fb96d95d31808c18887d1034eeb50c2048df01d3e12210806788f8aa4bd3888ff2c197a5126a0006a1a144bab24a8fc847b279e53b0564657cff91610a1dd32f823d508d13d7acdf27fc52a12755aeb3d42243c0aa22c00c062d12548bc4a0c3dbfe2bc8e809775b05393c051e7c738118b92573f0bacae19d83bf9834f1588285725b68fc35d9d41645d114aab9b6b4f1fd90e19975de1bc503d658416bb1e49b64777d53a74053b9b96d266cbb8c39d0a69ac95a83b81dcd1cde3b96df6683e98bfd8f11a49b48f64c91dd4b12387cd7799d330459edff38c0acbb3754848d6a07cd14fd646cab8c1dcf0f0b4c542c6a771d72aca6f0c29bc3606d9f5207fa257217df2e60cdc77a9e85af532d54eab6dd1aaddf585c46c3a903941d6fc063038c3a5c9486f944eb49dce4a3e002e553884522b6b151f7c1c1f37d91a15bb41c8320aa51f597744d2ca7e3ce9c90639de83d6b880fff4f1d56cd0e9480992a37827e4b6464bc38b4dce268287fc9f1f2b6efed4cbd11f32e52a8d40e3284170203e30c88b8483aadf99c299cd4e0865cb227166df1443b7d52870edc2587e0db66616cfc35e38029d741f02253c82dab5e1bbd25c5e029e582c0c71c507a1daa591057072f704a26bea31f72801b5029dfd9d9566c724cd5591ab305621d55ba94daea5b1712725963c58e196e68049108c80a95f7193bfe3b7f1dfdcc1de72eead76074dee6868bab6a32b7dd0758a85799628cd5c7f3be18459ca6f9537693df592af4e1f825d42392559089e8b42f3953e96a964f9c4fd5fa85f677cddc254f9a876cc2730bef01f02befe5d84a66ca46af751975579a6629f4e74c42927e5e2bff00c9b000ae5dd482efbd9717a7b79a183ca219a65477a1e2bde3386a8669d1bb149ffcc38b60b47108971d90cea564971f219816b99f592f2f230b5b86b9198a98abe982d2c0b0e5f40dd3697c23c6431b60088b2dc2a95f0c70214983d7a0015048e22fc3801b26119d38d435aafba92bb3b62b24d7fdfebac496e4e5f080b4e1f16ac2f4e0116ac7d9f16e2c945e7ce729f3e26612bcf37ce4c5bc34a553a5606e5f0624a61c0209bfa0a4c7ae3aa541b1ad71d9bf3fad478ca04bb53f3e7495b64c4951ffa530f4ccb1442fd5e350d581edcb9719cf3fc319fd9bca34a99c64d3f3ccda3f039522b459a7b95944e03dd3184611f82a88779a8c370390dd1039ecfa280d715ca8aa3dc5eacfbdbcc08ab4a579f62d6f2f9a7a67d96362c5a30245a6f412ba53ca6a276f632e765560e9e801949bf806f7557942a39700bd0dab24eecc510702a5f7fcfc7af70f37e040cbcb3d0bc6a2aa6107a2466685787ea9454e98d4b48789791017a6f3db37305e606dc12adabfc8d4f9d0651806e1ac2aafdad4762ec81487da70b2871eb8206b51abfb0b4694d563007fa000213182cea6ab976672f980300536258a409830a5f705537b93277250e1dddd77db1167cfdadea4ed4e6d0760b7aac94091b610aece5247b34c5ed042ca09a1a352014aad17a2141ac4c821ba4c8371d49cfd0008c3ef6c307e54f6decd2f6ee99d8de36a0a50d2f257919f30b48c8d22ac9ad0c"

	var tx, dup_tx transaction.Transaction

	// initialize globals
	//globals.Initialize()

	// setup logging
	globals.Logger = log.New()
	globals.Logger.SetLevel(log.DebugLevel)

	tx_raw, _ := hex.DecodeString(tx_hex)
	err := tx.DeserializeHeader(tx_raw)
	dup_tx.DeserializeHeader(tx_raw)

	if err != nil {
		t.Errorf("Tx Deserialisation failed")
	}

	pool, err := Init_Mempool(nil)

	if err != nil {
		t.Errorf("Pool initialization failed")
	}

	if len(pool.Mempool_List_TX()) != 0 {
		t.Errorf("Pool should be initialized in empty state")
	}

	if pool.Mempool_Add_TX(&tx, 0) != true {
		t.Errorf("Cannot Add transaction to pool in empty state")
	}

	if pool.Mempool_TX_Exist(tx.GetHash()) != true {
		t.Errorf("TX should already be in pool")
	}

	/*if len(pool.Mempool_List_TX()) != 1 {
		t.Errorf("Pool should  have 1 tx")
	}*/
	list_tx := pool.Mempool_List_TX()

	if len(list_tx) != 1 || list_tx[0] != tx.GetHash() {
		t.Errorf("Pool List tx failed")
	}

	get_tx := pool.Mempool_Get_TX(tx.GetHash())

	if tx.GetHash() != get_tx.GetHash() {
		t.Errorf("Pool get_tx failed")
	}

	// re-adding tx should faild
	if pool.Mempool_Add_TX(&tx, 0) == true || len(pool.Mempool_List_TX()) > 1 {
		t.Errorf("Pool should not allow duplicate TX")
	}

	// modify  tx and readd
	dup_tx.Unlock_Time = 99999999 //modify tx so txid changes, still it should be rejected

	if tx.GetHash() == dup_tx.GetHash() {

		t.Errorf("tx and duplicate tx must have different hash")
	}

	if pool.Mempool_Add_TX(&dup_tx, 0) == true || len(pool.Mempool_List_TX()) > 1 {
		t.Errorf("Pool should not allow duplicate Key images")

	}

	// pool must have  1 key_image

	key_image_count := 0
	pool.key_images.Range(func(k, value interface{}) bool {
		key_image_count++
		return true
	})

	if key_image_count != 1 {
		t.Errorf("Pool doesnot have necessary key image")
	}

	if pool.Mempool_Delete_TX(dup_tx.GetHash()) != nil {
		t.Errorf("non existing TX cannot be deleted\n")
	}

	// pool must have  1 key_image
	key_image_count = 0
	pool.key_images.Range(func(k, value interface{}) bool {
		key_image_count++
		return true
	})
	if key_image_count != 1 {
		t.Errorf("Pool must have necessary key image")
	}

	// lets delete
	if pool.Mempool_Delete_TX(tx.GetHash()) == nil {
		t.Errorf("existing TX cannot be deleted\n")
	}

	key_image_count = 0
	pool.key_images.Range(func(k, value interface{}) bool {
		key_image_count++
		return true
	})
	if key_image_count != 0 {
		t.Errorf("Pool should not have any key image")
	}

	if len(pool.Mempool_List_TX()) != 0 {
		t.Errorf("Pool should  have 0 tx")
	}

}
