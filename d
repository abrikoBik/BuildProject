[33mcommit 1a0796ec7c72fc68eec2ebb2474eef454c45f130[m[33m ([m[1;36mHEAD[m[33m -> [m[1;32mattachmentsNprojects[m[33m, [m[1;31morigin/main[m[33m, [m[1;31morigin/HEAD[m[33m, [m[1;32mmain[m[33m)[m
Merge: 3c1ceea 781a332
Author: e829dce <165459200+e829dce@users.noreply.github.com>
Date:   Sun Aug 24 16:36:06 2025 +0300

    Merge pull request #30 from abrikoBik/endpoints
    
    Endpoints and cleanup

[33mcommit 781a332745e37d733860e9ccc1c3c95a0712f8d6[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 16:32:21 2025 +0300

    models: clean up request and responses

[33mcommit 9d4e4004d7db5f77bc441ef76f5dd5b6e6819e2d[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 16:16:01 2025 +0300

    util: move constants to new package

[33mcommit 267bc1d9f279546529ca82a01fd2dce906684321[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 16:10:42 2025 +0300

    l10n: move errors from util

[33mcommit e2929601ceb824a6a15fd8db8988dec5efdc8eaa[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 15:57:13 2025 +0300

    l10n: organize into files

[33mcommit 7fe079007512692130999533b1fc14318d95cafa[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 15:38:01 2025 +0300

    api: move endpoints to a separate module

[33mcommit ec508ee818471bdbc53e633b0606db0206927c22[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 15:29:22 2025 +0300

    api: cleanup

[33mcommit 3c1ceeae428602f9b57906706c4cbd22daf3f343[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 14:54:24 2025 +0300

    db: convert to normal image, frontend: fixup

[33mcommit 96ee30d7ab27212726a830d79feaf543e14cd4d7[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 14:48:21 2025 +0300

    proxy: return trailing slashes

[33mcommit cd5061ea5b85dd7683fb0363fb0533a9c7aaaac9[m
Merge: 60c62fe a2a28a1
Author: e829dce <165459200+e829dce@users.noreply.github.com>
Date:   Sun Aug 24 00:18:08 2025 +0300

    Merge pull request #29 from abrikoBik/nginx
    
    Nginx

[33mcommit a2a28a1750959af3004a077e8fad3bf97ecb83d9[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 00:16:16 2025 +0300

    proxy: correct limit_except
    
    restrict frontend requests to requests that can be sent by browsers

[33mcommit a36a9ed7c01581605e33e813a2276e06756a201f[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 24 00:07:16 2025 +0300

    proxy: improve config
    
    - restrict request types
    - switch between release and dev frontend servers dynamically
    - remove trailing slashes
    - improve hot module reloading support

[33mcommit afcbf827eee1b9275f9135dca242b5a1df761e75[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 23:45:03 2025 +0300

    proxy: add container name

[33mcommit 4708d2206bac1596a52f6d9504e3637658d90ea6[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 23:43:58 2025 +0300

    all: utilize nginx

[33mcommit 2a8f8f3c1f0b3f563f0e25ea1a22a9aadda18e31[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 23:27:45 2025 +0300

    all: stop exposing backend and frontend ports

[33mcommit e8ef06101d9f7b847827b7d4f3353f44b1b44044[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 23:25:03 2025 +0300

    all: initial nginx support

[33mcommit 60c62fed4decc5a1516dbffcc952730aa3e72dc8[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 22:14:48 2025 +0300

    models: remove redundant attachment fields

[33mcommit e947c5e3f30b44b455d4173bd9498befd17d8372[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 22:10:27 2025 +0300

    frontend: improve dockerfile
    
    - use normal builder and runner images (curl ships out of the box)
    - copy lockfile for the builder

[33mcommit 6b98ba5906819f1ee3aad50740c9fb110dc9a542[m
Merge: ccf9186 0ff3189
Author: e829dce <165459200+e829dce@users.noreply.github.com>
Date:   Sat Aug 23 22:06:19 2025 +0300

    Merge pull request #23 from abrikoBik/attachments
    
    Attachments

[33mcommit 0ff318903d2dccffc30acd0b257d3fd023969a29[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 22:04:11 2025 +0300

    api: return attachment filename as well

[33mcommit d0626c072eb6028bf859fb3e0407e61c5d9bef4c[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 21:55:21 2025 +0300

    api: add uuid handling

[33mcommit 0e4cc3d736db9333706b43a0b9f308beaa022c5c[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 21:11:16 2025 +0300

    models: improve attachment creation logic
    
    must stop creating the attachment if the size is known to be over the limit

[33mcommit d649f54ef1750abdfffae880695df23bd9996f7c[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 21:09:23 2025 +0300

    backend: add docker volume for attachments

[33mcommit e1b468fad9ac53db36d12a0c7974ab58a3586f0b[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 20:55:28 2025 +0300

    backend: finalize attachment creation
    
    should now be concurrency-safe, since each goroutine now creates its own reader
    might not be as efficient though

[33mcommit 54c686d7d6f4d12bed0d0242f38688049a7d5788[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 23 20:21:53 2025 +0300

    models: add ProjectID property to attachments
    
    should be used as the foreign key, also fixes migrations

[33mcommit 0c622b6ec3373d2840de12703a38d77310f0e7c8[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 21:47:05 2025 +0300

    models: attachment.Size int -> int64

[33mcommit 6b1e36c1c9391e69675c152bc036d147c3344afc[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 21:45:49 2025 +0300

    cmd: add attachment to migrations

[33mcommit c236614ebe1f4810fe1e661a170e0d19edd6312c[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 21:44:14 2025 +0300

    models: simplify attachment creation && api: add error checks to storage creation

[33mcommit dcf0814f62c185eb2d4e142a1c8643e9f36bee13[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 21:23:21 2025 +0300

    api: create StoragePath directory on init

[33mcommit 7befa08023f2df33271fcc36673eb5d4d0012438[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 21:18:52 2025 +0300

    api: add multipart/form-data to supported content types

[33mcommit fd7ecace53ec9338d3f01fb3d7ce585c5bcc0cdb[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 20:59:01 2025 +0300

    api: add attachments endpoints

[33mcommit 788682797603636fbbdd90ecddb8d02762dde63a[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 20:54:17 2025 +0300

    chore: check for all errors

[33mcommit 861cd7e69ed0b4488f47b469548bcb55eb8df571[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 20:52:06 2025 +0300

    models: add attachment deletion

[33mcommit a475426ef8eb32e0060242c7d6d2c21a816b876b[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 20:48:29 2025 +0300

    models: parallelize attachment creation
    
    should be safe to do, since io.Copy does not modify src

[33mcommit e5bbfaab3b98c8d4057766856b6a2eb6d54432c0[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 01:15:03 2025 +0300

    chore: stop linting error strings

[33mcommit 42ca0e96dd8eb455744bf9fe47d1294ae90c0320[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 01:11:51 2025 +0300

    attachments: major refactor
    
    now attachment creation can be done in one request thanks to multipart

[33mcommit 61e49594cc381c658beebc17b2c34740d0f752b8[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 00:31:51 2025 +0300

    api: add attachment getter

[33mcommit 0cafe83ff0958e967d9525f6110860ffa947b70d[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 00:16:30 2025 +0300

    fixup! add empty function body
    
    for some reason, gopls was ok with that, but the builds were failing

[33mcommit a1bf8d967496bb22924aeca7ce25c2a32198880f[m
Author: e829dce <e829dce@disr.it>
Date:   Tue Aug 19 00:11:42 2025 +0300

    api: add attachment creation
    
    should work like this:
    1. receive json request with some metadata (currently just the name)
    2. create a database entry
    3. redirect to the endpoint where the actual attachment upload happens (`/attachments/<uuid>`)

[33mcommit ccf9186a35d08c52c50eb9fa5f4a22bac5304e89[m
Author: abereje <aelilkidae@gmail.com>
Date:   Mon Aug 18 17:39:17 2025 +0300

    Gantt chart was added

[33mcommit 70117218d2645a7fa7e46fc884099d42dd9bd0f2[m
Author: abereje <aelilkidae@gmail.com>
Date:   Mon Aug 18 15:35:52 2025 +0300

    Made blueprints scalable and movable

[33mcommit c9204e05c6241af0bc489f5872f252f0bffe8376[m
Author: abereje <aelilkidae@gmail.com>
Date:   Sun Aug 17 17:50:41 2025 +0300

    Fix backend communication for Windows

[33mcommit 6d61898586fa4f14ba2c87ddc5df0530a9c16638[m
Author: e829dce <e829dce@disr.it>
Date:   Sun Aug 17 01:28:36 2025 +0300

    attachments: create model
    
    also implemented get, create and delete; patch not planned, as it is going to be implemented in the frontend

[33mcommit d1ac3cf61d63f928f0ad6d268f25b633bdd6270c[m
Author: e829dce <e829dce@disr.it>
Date:   Sat Aug 16 18:28:17 2025 +0300

    chore: bump modules

[33mcommit f1e6f181ff0115a8b475278713a978ec7bbd6e1a[m
Author: e829dce <e829dce@disr.it>
Date:   Fri Aug 15 22:50:42 2025 +0300

    models: fix worker-stage relationship; backend: lint

[33mcommit 73bd7f4863b41a28f52fd05d77d084bb56940085[m
Author: abereje <aelilkidae@gmail.com>
Date:   Thu Aug 14 19:18:07 2025 +0300

    Remove Gantt chart for now
    
    - Fix npm trying to install legacy packages

[33mcommit 43c1019cd84d4f0f6713d811c33cb60d1d476a63[m
Author: abereje <aelilkidae@gmail.com>
Date:   Mon Aug 11 17:45:36 2025 +0300

    Gantt chart was added and slight fixes on PushNotification.tsx

[33mcommit caabaa6cd4348d281a7ee490247c9bc23ec80398[m
Author: abereje <aelilkidae@gmail.com>
Date:   Mon Aug 11 16:02:22 2025 +0300

    Fix diagram animation dont works
    
    - Linting have 1 warning but I think this okay

[33mcommit 467fea3d3a383d97fdafdcaa724366c2098e9e5a[m
Author: abereje <aelilkidae@gmail.com>
Date:   Sat Aug 9 22:20:12 2025 +0300

    Fix linting errors

[33mcommit becdf1946f843344d4b60c0bf5f6ece53337fd11[m
Author: abereje <aelilkidae@gmail.com>
Date:   Sat Aug 9 15:59:36 2025 +0300

    Smooth animation to the chart on main page

[33mcommit b66211fa1304cb6ae59bd1bf5805dca8d66233fd[m
Author: abereje <aelilkidae@gmail.com>
Date:   Sat Aug 9 15:05:21 2025 +0300

    Sligt fixes
    
    - Fix links on homepage
    - Fix zoom function on blueprints page
    - Fix edit panel margins

[33mcommit b713b088f60e5115eb08b28ea8f902a349c1b4d1[m
Author: abereje <aelilkidae@gmail.com>
Date:   Sat Aug 9 14:41:31 2025 +0300

    Fix typescript linting errors

[33mcommit 182626df9d0644c5b4be9430ab447fa074b35fb6[m
Author: abereje <aelilkidae@gmail.com>
Date:   Sat Aug 9 14:36:31 2025 +0300

    Now blueprints are editable
    
    - Blueprints are editable
    - Gantt diagram placeholder was added

[33mcommit 93a6a68c2fc231337e524c20dbbb9960feb54868[m
Author: abereje <aelilkidae@gmail.com>
Date:   Fri Aug 8 14:36:32 2025 +0300

    Add some features on Estimates.tsx and some fixes
    
    - Estimate lines are now deletable and editable
    - Moved push notification to independent component (little bit buggy rn)
    - Fix react warning for CookiesProvider

[33mcommit 11e55990dc6022cb391a588c55efe8485584641c[m
Author: abereje <aelilkidae@gmail.com>
Date:   Thu Aug 7 14:38:19 2025 +0300

    Fix typescript linting errors

[33mcommit b81774c83cbafb700b57175db02d3eaa3fccd28e[m
Author: abereje <aelilkidae@gmail.com>
Date:   Wed Aug 6 20:49:19 2025 +0300

    Update blueprints.tsx estimates.tsx and home.tsx pages                                                                                                                                                                                          - Now you can zoom blueprint (for now you cant move it and zoom more)                                                   - At estimates you can add new lines for estimate                                                                       - Number and email now copiable                                                                                         - Push notification added (not in independed component)

[33mcommit 5f8e6fdb776a25fad63c729fbfaf877fc84a241a[m
Merge: a4ee85d a061162
Author: Daniil Popov <abereje@localhost.localdomain>
Date:   Sat May 31 00:03:28 2025 +0300

    Merge branch 'main' of github.com:abrikoBik/BuildProject

[33mcommit a4ee85d340f54b8ca99fb7533e81b5d9e1939269[m
Author: Daniil Popov <abereje@localhost.localdomain>
Date:   Sat May 31 00:01:07 2025 +0300

    Add buttons on Home.tsx
    
    - Problems with adaptive, cus of buttons use some percentage of parent
      div (Not so urgent but need to fix as soon as posible)

[33mcommit b1f1390c90f27478a4a84d29346cfb2d4f15eed1[m
Author: Daniil Popov <abereje@localhost.localdomain>
Date:   Tue May 27 16:35:31 2025 +0300

    Some bug fixes
    
    - Now you cant select icon ;))))))))

[33mcommit a061162bda27014207e41e2f6f300c49d98d6cee[m
Merge: 71d6458 c51596e
Author: e829dce <165459200+e829dce@users.noreply.github.com>
Date:   Tue May 27 14:11:29 2025 +0300

    Merge pull request #19 from abrikoBik/stages
    
    Stages

[33mcommit c51596e1c79c4e3100f9370095a327033a46b499[m
Author: e829dce <e829dce@disr.it>
Date:   Tue May 27 14:09:37 2025 +0300

    models: finalize project, group, stage

[33mcommit 71d6458c872fdee214c138f16d8f2d74df241d52[m
Author: Daniil Popov <abereje@localhost.localdomain>
Date:   Tue May 27 01:27:08 2025 +0300

    Made up all pages (they not functional for now)
    
    - Sidebar hide is bugged when on narrow screen
    - Buttons dont usable at all, they dont interact with user

[33mcommit 11e4e17861de495728eaaec11ca55a4908a92fc1[m
Author: e829dce <e829dce@disr.it>
Date:   Mon May 26 13:53:32 2025 +0300

    models, api: finalize workers
    
    creating/setting works properly now

[33mcommit 65d68f495dca2f8ae1cdbb6e850d534095de110a[m
Author: e829dce <e829dce@disr.it>
Date:   Mon May 26 01:27:12 2025 +0300

    works done for today
    
    - fix worker registration
    - more verbose migration
    - fix relationships

[33mcommit 16c5833954c82d6e8f1c08ac44b7e58456773b4c[m
Author: e829dce <e829dce@disr.it>
Date:   Sun May 25 19:45:31 2025 +0300

    models: add project methods

[33mcommit d4e9c74bed9c34dd63d232c0542eed68c1202c84[m
Author: e829dce <e829dce@disr.it>
Date:   Sun May 25 19:15:53 2025 +0300

    models: correct project model

[33mcommit 9af5e6767ecaf8ae2b2b0c31edb70cae7073989e[m
Merge: c163f25 a69221f
Author: e829dce <165459200+e829dce@users.noreply.github.com>
Date:   Sun May 25 19:07:09 2025 +0300

    Merge pull request #20 from abrikoBik/openapi
    
    OpenAPI specification

[33mcommit a69221f24af2f56dcd65625e641c42f6e6703351[m[33m ([m[1;31morigin/openapi[m[33m)[m
Author: e829dce <e829dce@disr.it>
Date:   Sun May 25 19:06:22 2025 +0300

    i foror

[33mcommit 51016d51bea630926771cd9616cb8997ae1d9d7f[m
Author: e829dce <e829dce@disr.it>
Date:   Sat May 24 21:14:52 2025 +0300

    spec: The Great Split

[33mcommit cc711bef8d65ff7a0c092491ae162e3bc5eb6241[m
Author: e829dce <e829dce@disr.it>
Date:   Sat May 24 20:56:00 2025 +0300

    spec: add responses

[33mcommit 369c5ad2fc415319671ae08b14a358b0369112c8[m
Author: e829dce <e829dce@disr.it>
Date:   Sat May 24 20:26:33 2025 +0300

    spec: initial spec

[33mcommit 2c9034831ebc016995fdb832b3a0c1c535dfa0d9[m
Author: e829dce <e829dce@disr.it>
Date:   Sat May 24 19:25:08 2025 +0300

    api: fix cors

[33mcommit f7926ba2cfa4e90b9431710007af86dfd1e37185[m
Author: e829dce <e829dce@disr.it>
Date:   Sat May 24 18:50:21 2025 +0300

    api: add GetStages

[33mcommit 3435d20f23a4a9e892192d1620a210cd9ad05a72[m
Author: e829dce <e829dce@disr.it>
Date:   Sun May 18 16:04:55 2025 +0300

    models: add stage methods

[33mcommit c163f25a84747e63dae939e55c71e4bb4adcc10e[m
Merge: cb69101 98692b3
Author: e829dce <165459200+e829dce@users.noreply.github.com>
Date:   Tue May 13 14:44:53 2025 +0300

    Merge pull request #18 from abrikoBik/worker-endpoints
    
    Worker endpoints

[33mcommit 98692b372dd5e4bf8224d37087ed1849561e2c94[m
Author: e829dce <e829dce@disr.it>
Date:   Tue May 13 14:39:36 2025 +0300

    worker: add GetWorker method
    
    untested
