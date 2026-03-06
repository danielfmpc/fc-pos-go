---
description: Regras para gerar mensagens de commit
---
Sempre que for solicitado a sugerir ou gerar uma mensagem de commit (mesmo implicitamente):

1. **Formato**: Use o formato Conventional Commits: `<tipo>(<escopo>): <descrição breve>`.
2. **Tipos válidos**: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`.
3. **Escopo**: Use o nome do módulo, componente ou arquivo principal sendo alterado (opcional, mas recomendado).
4. **Descrição (Subject)**:
   - Comece com verbo conjugado na terceira pessoa do singular ou no imperativo afirmativo (ex.: `Adiciona`, `Corrige`, `Remove`, `Atualiza` ou `Add`, `Fix`, etc., dependendo do idioma do repositório).
   - Máximo de 50 caracteres.
   - Não coloque ponto final.
5. **Corpo do Commit (Body)**:
   - Para mudanças que não sejam triviais, adicione uma linha em branco após a descrição e inclua um corpo explicando o **QUÊ** e o **POR QUÊ** da mudança (evite focar no 'como', pois o código já mostra isso).
   - Quebre as linhas do corpo em 72 caracteres.
6. **Qualidade**: Evite descrições genéricas como `Atualiza arquivos` ou `Corrige bug`. Seja específico e claro sobre qual é o valor adicionado na alteração.
